package adstxt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/mediaFORGE/supplyqc/adpath/infra/data/enum"
)

var Client = http.Client{

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
	l := d.DomainList()

	// choose the ads.txt with the most specific url as long as it is not also provided by a less specific url.

	for _, d := range l {
		url, data, err := Get(d)
		if err != nil {
			errs = fmt.Errorf("%s error encountered while getting %s, %s", errs, d, err)
		}
	}

}

func Get(domain string) (url string, bytes []byte, err error) {
	resp, err := Request(domain)
	if err != nil {
		return "", nil, err
	}
	b, err := Read(resp)
	if err != nil {
		return "", nil, err
	}
	return resp.Request.URL.String(), b, err

}

func Request(domain string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://" + domain + "/ads.txt", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept","text/plain" )
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

