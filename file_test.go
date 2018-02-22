package adstxt

import (
	"testing"

	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/stretchr/testify/assert"
)

var rec0 = Record{
	AdSystemDomain:        "advertising.amazon.com",
	CanonicalSystemDomain: "amazon-adsystem.com",
	PublisherID:           "pub-123",
	AccountType:           DIRECT,
	CertAuthorityID:       "qwerty",
	Ext:                   []string{"test1", "test2"},
	Comment:               "stuff",
	LineNum:               1, // 0 is the first line
}

var rec0string = "advertising.amazon.com, pub-123, DIRECT, qwerty; test1;test2 # stuff "

var variable = Variable{
	Key:     "contact",
	Value:   "adops@example.com",
	LineNum: 0,
}

var variable0string = " contact = adops@example.com "

func TestLineComment(t *testing.T) {
	s := "# a line comment\n" + rec0string + "\n# another line comment"
	rec, lc, va, el, e := Parse([]byte(s))
	assert.Exactly(t, []Record{rec0}, rec)
	assert.Nil(t, va)
	assert.Nil(t, el)
	assert.Nil(t, e)
	assert.Exactly(t, []LineComment{
		{Text: "# a line comment", LineNum: 0},
		{Text: "# another line comment", LineNum: 2},
	}, lc)

	s = rec0string
	_, lc, _, _, e = Parse([]byte(s))
	assert.Nil(t, e)
	assert.Nil(t, lc)
}

func TestParseVariable(t *testing.T) {
	v, err := ParseVariable(variable0string)
	assert.Nil(t, err)
	assert.Exactly(t, variable, *v)

	v, err = ParseVariable("")
	assert.NotNil(t, err)
	assert.Nil(t, v)
	v, err = ParseVariable("this=that=other")
	assert.NotNil(t, err)
	assert.Nil(t, v)
}

func TestParseRecord(t *testing.T) {
	r, err := ParseRecord(rec0string)
	assert.Nil(t, err)
	rec0.LineNum = 0
	assert.Exactly(t, rec0, *r)

	r, err = ParseRecord("#comment")
	assert.Error(t, err)
	assert.Nil(t, r)

	r, err = ParseRecord("google.com, pubid123 #this is invalid")
	assert.Error(t, err)
	assert.Nil(t, r)

	r, err = ParseRecord("google.com, pubid123, OTHERTYPE #this is invalid")
	assert.Error(t, err)
	assert.Nil(t, r)

	r, err = ParseRecord("google.com, pubid123, DIRECT, trustme123, invalid5thfield")
	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestParse(t *testing.T) {
	// filter html
	r, l, v, e, err := Parse([]byte("<html> google.com, pubid123, DIRECT, trustme123 </html>"))
	assert.Error(t, err)
	assert.Nil(t, r)
	assert.Nil(t, l)
	assert.Nil(t, v)
	assert.Nil(t, e)

	// good variable
	r, l, v, e, err = Parse([]byte(" contact = adops@example.com \n" + rec0string))
	assert.Nil(t, err)
	rec0.LineNum = 1
	assert.Exactly(t, []Record{rec0}, r)
	assert.Nil(t, l)
	assert.Exactly(t, []Variable{variable}, v)
	assert.Nil(t, e)

	// bad variable
	r, l, v, e, err = Parse([]byte(" = adops@example.com \n" + rec0string))
	assert.Nil(t, err)
	rec0.LineNum = 1
	assert.Exactly(t, []Record{rec0}, r)
	assert.Nil(t, l)
	assert.Nil(t, v)
	assert.NotNil(t, e)

	// valid lines but no valid adstxt records
	r, l, v, e, err = Parse([]byte("contact = adops@example.com \n"))
	assert.Error(t, err)
	assert.Nil(t, r)
	assert.Nil(t, l)
	assert.Nil(t, v)
	assert.Nil(t, e)

	// bad record produces error line
	r, l, v, e, err = Parse([]byte(rec0string + "\nadvertising.amazon.com, pub-123, OTHER, qwerty; test1;test2 # stuff"))
	assert.Nil(t, err)
	rec0.LineNum = 0
	assert.Exactly(t, []Record{rec0}, r)
	assert.Nil(t, l)
	assert.Nil(t, v)
	assert.NotNil(t, e)
	rec0.LineNum = 1
}

func TestNewFile(t *testing.T) {
	// nil bytes
	f, err := NewFile([]byte{}, time.Now(), "http://www.test.com", "test.com", "test.com")
	assert.Nil(t, f)
	assert.Error(t, err)

	// zero time
	f, err = NewFile([]byte(rec0string+"\n"+variable0string), time.Time{}, "http://www.test.com", "test.com", "test.com")
	assert.Nil(t, f)
	assert.Error(t, err)

	// no url
	f, err = NewFile([]byte(rec0string+"\n"+variable0string), time.Now(), "", "test.com", "test.com")
	assert.Nil(t, f)
	assert.Error(t, err)

	// no root url
	f, err = NewFile([]byte(rec0string+"\n"+variable0string), time.Now(), "http://www.test.com", "", "test.com")
	assert.Nil(t, f)
	assert.Error(t, err)

	// no adstxt url
	f, err = NewFile([]byte(rec0string+"\n"+variable0string), time.Now(), "http://www.test.com", "test.com", "")
	assert.Nil(t, f)
	assert.Error(t, err)

	// failed parse
	f, err = NewFile([]byte("invalid adstxt text file"), time.Now(), "http://www.test.com", "test.com", "test.com")
	assert.Nil(t, f)
	assert.Error(t, err)

	//successful parse
	f, err = NewFile([]byte(variable0string+"\n"+rec0string), time.Now(), "http://www.test.com", "test.com", "test.com")
	assert.Nil(t, err)

	sh := sha256.Sum256([]byte(variable0string + "\n" + rec0string))
	cs := base64.StdEncoding.EncodeToString(sh[:])

	file := &File{
		URL:          "http://www.test.com",
		RootDomain:   "test.com",
		AdstxtDomain: "test.com",
		Records:      []Record{rec0},
		Variables:    []Variable{variable},
		CheckSum:     cs,
	}

	assert.Exactly(t, file, f)

}