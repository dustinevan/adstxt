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

func Crawl(domain string) (file []byte, url string, err error) {
	d, err := DomainFromString(domain)
	if err != nil {
		return nil, "", err
	}

	var errs error
	l := d.ListDomains()

	// start with the root domain, if that domain returns an adstxt the contains a subdomain variable that
	// matches any of the listed domains, travel there recursively. Keep the last valid ads.txt. If ads.txt
	// files match its an error, keep the last valid ads.txt found.
	var f *File
	for _, dom := range l {
		host, data, err := Get(dom)
		if err != nil {
			errs = fmt.Errorf("%s error encountered while getting %s, %s", errs, dom, err)
			continue
		}
		newf, err := NewFile(data, time.Now(), dom, l[0], host)
		if err != nil {
			errs = fmt.Errorf( "%s error encountered while parsing adstxt from %s, %s", errs, dom, err)
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



	}

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
