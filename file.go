package adstxt

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"unicode"

	"github.com/dustinevan/chron"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

type File struct {
	// The URL location of this adstxt
	URL string `json:"url"`

	// The Root Domain
	RootDomain string `json:"root_domain"`

	// The Subdomain of the Root Domain the adstxt is valid for
	AdstxtDomain string `json:"adstxt_domain"`

	// Valid Exchange/PubID combinations/routes for a certain publishers bid requests
	Records []Record `json:"adpaths"`

	// Comments that occupy a full line
	LineComments []LineComment `json:"line_comments,omitempty"`

	// Any line containing a pattern of <VARIABLE>=<VALUE> should be interpreted as a variable
	// declaration.
	Variables []Variable `json:"variables,omitempty"`

	//
	ErrLines []string

	// SHA256 checksum of the bytes in the response body
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

	cs := sha256.Sum256(b)

	file = &File{
		URL:          url,
		RootDomain:   root,
		AdstxtDomain: adstxtdom,
		Records:      make([]Record, 0),
		LineComments: make([]LineComment, 0),
		Variables:    make([]Variable, 0),
		CheckSum:     string(cs[:]),

	}

	//unparsedlines, errs = file.parse(b)

	return
}

func Parse(b []byte) (rec []Record, lc []LineComment, va []Variable, el []ErrorLine, e error) {
	if strings.Contains(string(b), "<html") {
		return nil, nil, nil, nil, fmt.Errorf("parser encountered html")
	}

	lines := removeEmptyLines(string(b))

	for i, line := range lines {
		if line == "" {
			panic("empty line, this is a bug")
		}

		if line[0] == '#' {
			lc = append(lc, LineComment{Text: line, LineNum: i})
			continue
		}

		if strings.Contains(line, "=") {
			v, err := ParseVariable(line)
			if err != nil {
				el = append(el, ErrorLine{Error: err, Line: line, LineNum: i})
				continue
			}
			v.LineNum = i
			va = append(va, *v)
			continue
		}

		r, err := ParseRecord(line)
		if err != nil {
			el = append(el, ErrorLine{Error: err, Line: line, LineNum: i})
			continue
		}
		rec = append(rec, *r)
	}

	if len(rec) == 0 {
		return nil, nil, nil, nil, fmt.Errorf("parser found no valid adstxt records")
	}

	return
}

func ParseRecord(line string) (*Record, error) {
	r := &Record{}

	// get comments
	i := strings.Index(line, "#")
	if i > -1 {
		if i == 0 {
			return nil, fmt.Errorf("this is a line comment")
		}
		r.Comment = line[i:]
		line = line[:i]
	}

	// get extensions
	if strings.Contains(line, ";") {
		s := strings.Split(line, ";")
		line = s[0]
		r.Ext = s[1:]
	}

	// split fields
	line = strings.Trim(line, ",")
	fields := strings.Split(line, ",")
	if len(fields) < 3 {
		return nil, fmt.Errorf("%s only has %v fields, at least 3 required", line, len(fields))
	}

	r.AdSystemDomain = fields[0]

	// attempt to find a canonical ad system domain, do nothing if we can't find one
	adsys, err := GetCanonicalAdSystemDomain(fields[0])
	if err == nil {
		r.CanonicalSystemDomain = strings.ToLower(adsys)
	}

	r.PublisherID = fields[1]

	at := GetAccountType(fields[2])
	if at == INVALID_ACCOUNT_TYPE {
		return nil, fmt.Errorf( "encountered invalid account type %s", fields[2])
	}
	r.AccountType = at

	if len(fields) == 4 {
		r.CertAuthorityID = fields[3]
	}

	if len(fields) > 4 {
		return nil, fmt.Errorf("too many fields. found %v, expected 3 or 4", len(fields))
	}

	return r, nil
}

func ParseVariable(line string) (*Variable, error) {
	parts := strings.Split(line, "=")
	if len(parts) > 2 {
		return nil, fmt.Errorf("found too many parts while parsing variable")
	}
	if len(parts) < 2 {
		return nil, fmt.Errorf("no = found, this is not an adstxt variable")
	}
	return &Variable{ Key: parts[0], Value: parts[1] }, nil
}


func removeEmptyLines(file string) (lines []string) {
	all := strings.Split(file, "\n")
	for _, l := range all {
		line := removeWhiteSpace(l)
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	return lines
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
	// attaching line comment information to records.
	LineNum int `json:"line_num"`
}

type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Line number the comment was found on after removing empty lines. This is useful for
	// attaching line comment information to records.
	LineNum int `json:"line_num"`
}

type ErrorLine struct {
	// Reason the parse failed
	Error error `json:"error"`

	// Original data
	Line string `json:"line"`

	// Line number of the parse failure
	LineNum int `json:"line_num"`
}

type PublisherAccountType int

const (
	NO_ACCOUNT_TYPE_SPECIFIED PublisherAccountType = iota
	DIRECT
	RESELLER
	BOTH // some ads.txt file contain duplicate rows for the same pubid with reseller and direct types. these can be reduced by calling DedupOnAccountType()
	INVALID_ACCOUNT_TYPE
)

func GetAccountType(s string) PublisherAccountType {
	s = strings.ToUpper(s)
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
