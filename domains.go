package adstxt

import (
	"net/url"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"strings"
	"github.com/gin-gonic/gin/json"
)

type Domain struct {
	Host string // sub2.sub1.test.co.jp
	Root string // test.co.jp
	PublicSuffix string // co.jp
	ICANN bool // see PublicSuffix comments
	Subs []string // [ "sub2" "sub1" ] most specific first
}

func DomainFromURL(u *url.URL) (*Domain, error) {
	host := u.Hostname()
	if host == "" {
		return nil, fmt.Errorf("url.Hostname is blank")
	}

	root, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return nil, err
	}

	ps, icann := publicsuffix.PublicSuffix(host)
	if err != nil {
		return nil, err
	}

	s := strings.Replace(host, root, "", 1)
	s = strings.TrimRight(s, ".")
	subs := strings.Split(s, ".")

	return &Domain{
		Host: host,
		Root: root,
		PublicSuffix: ps,
		ICANN: icann,
		Subs: subs,
	}, nil
}

func DomainFromString(domain string) (*Domain, error) {
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	return DomainFromURL(u)
}

func (d *Domain) DomainList() []string {
	l := make([]string, 0)

	for i := 0; i < len(d.Subs); i++ {
		t := strings.Join(d.Subs[i:], ".")
		l = append(l, t + "." + d.Root)
	}
	l = append(l, d.Root)
	return l
}

func (d *Domain) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("%s|%s|%s|%s|%s", d.Host, d.Root, d.PublicSuffix, d.ICANN, d.Subs)
	}
	return string(b)
}