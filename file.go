package adstxt

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/dustinevan/chron"
	"github.com/json-iterator/go"
	"unicode"
	"github.com/rekrt/app-server/mon"
)

var json = jsoniter.ConfigFastest

type File struct {
	// The URL location of this adstxt
	URL string `json:"url"`

	// The Root Domain
	RootDomain string `json:"root_domain"`

	AdstxtDomain string `json:"adstxt_domain"`

	// Valid Exchange/PubID combinations/routes for a certain publishers bid requests
	Records []Record `json:"adpaths"`

	// Comments that occupy a full line
	LineComments []LineComment `json:"line_comments,omitempty"`

	// Any line containing a pattern of <VARIABLE>=<VALUE> should be interpreted as a variable
	// declaration.
	Variables []Variable `json:"variables,omitempty"`

	// SHA256 checksum of the ads.txt bytes
	CheckSum string `json:"checksum"`

	// The time of the adstxt get request
	LookupTime time.Time `json:"lookup_time"`
}

func NewFile(b []byte, t time.Time, url, root, adstxtdom string) (file *File, unparsedlines []string, errs []error) {
	if len(b) == 0 {
		return nil, nil, []error{fmt.Errorf("empty bytes passed")}
	}
	if t == chron.ZeroValue().AsTime() {
		return nil, nil, []error{fmt.Errorf("invalid lookup time passed")}
	}
	if url == "" || root == "" || adstxtdom == "" {
		return nil, nil, []error{fmt.Errorf("invalid urls passed, %s %s %s", url, root, adstxtdom)}
	}

	mon.EMAIL_SYSTEM

	file = &File{
		URL:          url,
		RootDomain:   root,
		AdstxtDomain: adstxtdom,
		Records:      make([]Record, 0),
		LineComments: make([]LineComment, 0),
		Variables:    make([]Variable, 0),
		CheckSum:     string(sha256.Sum256(b)[:]),
	}

	unparsedlines, errs = file.parse(b)

	return
}

func (f *File) parse(b []byte) (unparsedlines []string, err []error) {
	lines := strings.Split(strings.Replace(string(b), "\n\n", "\n", -1), "\n")

	for i, line := range lines {
		line = removeWhiteSpace(line)
		if line[0] == '#' && len(line) > 0 {

		}
	}

}

func removeWhiteSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, s)
}

func (f *File) String() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return fmt.Sprint(f)
	}
	return string(bytes)
}

type Record struct {
	// (Required) The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that
	// bidders connect to. This may be the operational domain of the system, if that is different than the
	// parent corporate domain, to facilitate WHOIS and reverse IP lookups to establish clear ownership of
	// the delegate system. Ideally the SSP or Exchange publishes a document detailing what domain name
	// to use.
	AdSystemDomain string `json:"ad_system_domain"`

	// This field is an attempt to reconcile different ad system domains that mean the same thing. Matching
	// adstxt data with bid request data requires a mapping, but because many adstxt files say the same thing
	// different this field attempts to canonize a specific ad system spelling see disambiguation.go
	CanonicalSystemDomain string `json:"canonical_system_domain"`

	// (Required) The identifier associated with the seller or reseller account within the advertising system in
	// field #1. This must contain the same value used in transactions (i.e. OpenRTB bid requests) in the
	// field specified by the SSP/exchange. Typically, in OpenRTB, this is publisher.id. For OpenDirect it is
	// typically the publisher’s organization ID. ExDomain.
	PublisherID string `json:"publisher_id"`

	// (Required) An enumeration of the type of account. A value of ‘DIRECT’ indicates that the Publisher
	// (content owner) directly controls the account indicated in field #2 on the system in field #1. This
	// tends to mean a direct business contract between the Publisher and the advertising system. A value
	// of ‘RESELLER’ indicates that the Publisher has authorized another entity to control the account
	// indicated in field #2 and resell their ad space via the system in field #1. Other types may be added
	// in the future. Note that this field should be treated as case insensitive when interpreting the data.
	AccountType PublisherAccountType `json:"account_type"`

	// (Optional) An ID that uniquely identifies the advertising system within a certification authority
	// (this ID maps to the entity listed in field #1). A current certification authority is the Trustworthy
	// Accountability Group (aka TAG), and the TAGID would be included here.
	CertAuthorityID string `json:"cert_authority_id,omitempty"`

	// Extension fields are allowed by implementers and their consumers as long as they utilize a
	// distinct final separator field ";" before adding extension data to each record
	Ext []string `json:"ext,omitempty"`

	// Anything after # on a line is considered to be a comment
	Comment string `json:"comment,omitempty"`

	// The line number the record was found on, after removing empty lines. This is useful for
	// attaching line comment information to records.
	LineNum int `json:"line_num"`
}

type LineComment struct {
	// Comment text
	Text string `json:"text"`

	// Line number the comment was found on after removing empty lines. This is useful for
	// attaching line comment information to records
	LineNum int `json:"line_num"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Line number the comment was found on after removing empty lines. This is useful for
	// attaching line comment information to records
	LineNum int `json:"line_num"`
}

type PublisherAccountType int

const (
	NO_ACCOUNT_TYPE_SPECIFIED PublisherAccountType = iota
	DIRECT
	RESELLER
	BOTH // optionally duplicate adsystem pubid combination with reseller and direct can be reduced to this
	INVALID_ACCOUNT_TYPE
)

func GetAccountType(s string) PublisherAccountType {
	switch s {
	case "":
		return NO_ACCOUNT_TYPE_SPECIFIED
	case "DIRECT":
		return DIRECT
	case "RESELLER":
		return RESELLER
	default:
		return INVALID_ACCOUNT_TYPE
	}
}
