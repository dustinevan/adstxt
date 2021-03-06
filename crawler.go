package adstxt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Client = http.Client{

	// TODO fix using base domain instead of host
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 1 {
			hosts := make(map[string]struct{})
			hosts[strings.Replace(req.URL.Host, "www.", "", 1)] = struct{}{}
			for _, r := range via {
				hosts[strings.Replace(r.URL.Host, "www.", "", 1)] = struct{}{}
			}
			if len(hosts) > 2 {
				return fmt.Errorf("illegal redirect chain. %s", hosts)
			}
		}
		return nil
	},
	Timeout: time.Second * 10,
}

func Crawl(domain string) (*File, error) {
	d, err := DomainFromString(domain)
	if err != nil {
		return nil, err
	}

	var errs []error
	l := d.ListDomains()

	// start with the root domain, if that domain returns an adstxt the contains a subdomain variable that
	// matches any of the listed domains, travel there recursively. Keep the last valid ads.txt. If ads.txt
	// files match its an error, keep the last valid ads.txt found.
	var f *File
	for _, dom := range l {
		// only travel to subdomain if it's given as a var in the parent
		if f != nil && !f.IsValidSubDomain(dom) {
			continue
		}
		host, data, err := Get(dom)
		if err != nil {
			errs = append(errs, fmt.Errorf("error encountered while getting %s, %s", dom, err))
			continue
		}
		newf, err := NewFile(data, time.Now(), dom, l[0], host)
		if err != nil {
			errs = append(errs, fmt.Errorf( "error encountered while parsing adstxt from %s, %s", dom, err))
			continue
		}
		// first adstxt found
		if f == nil {
			f = newf
			continue
		}
		// the subdomain has the same adstxt
		if f.CheckSum == newf.CheckSum {
			continue
		}
		f = newf
	}

	if f == nil {
		return nil, ErrJoin(errs, "|")
	}
	f.LookupTime = time.Now()
	return f, nil
}

func Get(domain string) (host string, bytes []byte, err error) {
	resp, err := Request(domain)
	if err != nil {
		return "", nil, err
	}
	b, err := Read(resp)
	if err != nil {
		return "", nil, err
	}
	return resp.Request.URL.Host, b, err

}

func Request(domain string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://"+domain+"/ads.txt", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/plain")
	return Client.Do(req)
}

func Read(resp *http.Response) ([]byte, error) {

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("request to %s returned status %s", resp.Request.URL, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error: %s when reading the response body from %s", err, resp.Request.URL)

	}
	if len(body) == 0 {
		return nil, fmt.Errorf("empty response body recieved from %s", resp.Request.URL)
	}

	return body, nil
}

func ErrJoin(errs []error, delim string) error {
	s := make([]string, 0)
	for _, e := range errs {
		s = append(s, e.Error())
	}
	return fmt.Errorf("%s", strings.Join(s, delim))
}