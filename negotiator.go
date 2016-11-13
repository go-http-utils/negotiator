package negotiator

import (
	"net/http"
	"strings"
)

const (
	// HeaderAccept is the HTTP "Accept" Header.
	HeaderAccept = "Accept"
	// HeaderAcceptLanguage is the HTTP "Accept-Language" Header.
	HeaderAcceptLanguage = "Accept-Language"
	// HeaderAcceptEncoding is the HTTP "Accept-Encoding" Header.
	HeaderAcceptEncoding = "Accept-Encoding"
	// HeaderAcceptCharset is the HTTP "Accept-Charset" Header.
	HeaderAcceptCharset = "Accept-Charset"
)

type spec struct {
	val string
	q   float64
}

// Specs it the shorthand for []Spec.
type specs []spec

// Len is used to impelement sort.Interface for Specs.
func (ss specs) Len() int {
	return len(ss)
}

// Swap is used to impelement sort.Interface for Specs.
func (ss specs) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

// Less is used to impelement sort.Interface for Specs.
func (ss specs) Less(i, j int) bool {
	if ss[i].q > ss[j].q {
		return true
	}
	if ss[i].q == ss[j].q {
		if ss[i].val == "*" || strings.HasSuffix(ss[i].val, "/*") {
			return true
		}

		if ss[j].val == "*" || strings.HasSuffix(ss[j].val, "/*") {
			return false
		}

		return i < j
	}

	return false
}

func (ss specs) hasVal(val string) bool {
	for _, spec := range ss {
		if spec.val == val {
			return true
		}
	}

	return false
}

func formatHeaderVal(val string) string {
	return strings.ToLower(strings.Replace(val, " ", "", -1))
}

// Negotiator repensents the HTTP negotiator.
type Negotiator struct {
	req *http.Request
}

// New creates an instance of Negotiator.
func New(req *http.Request) Negotiator {
	return Negotiator{req}
}

// Accept returns the most preferred content types from the HTTP Accept header.
func (n Negotiator) Accept(offers []string) (bestOffer string, matched bool) {
	parser := newHeaderParser(n.req.Header, true)

	return parser.selectOffer(offers, parser.parse(HeaderAccept))
}

// Language returns the most preferred language from the HTTP Accept-Language
// header.
func (n Negotiator) Language(offers []string) (bestOffer string, matched bool) {
	parser := newHeaderParser(n.req.Header, false)

	return parser.selectOffer(offers, parser.parse(HeaderAcceptLanguage))
}

// Encoding returns the most preferred language from the HTTP Accept-Encoding
// header.
func (n Negotiator) Encoding(offers []string) (bestOffer string, matched bool) {
	parser := newHeaderParser(n.req.Header, false)

	return parser.selectOffer(offers, parser.parse(HeaderAcceptEncoding))
}

// Charset returns the most preferred language from the HTTP Accept-Charset
// header.
func (n Negotiator) Charset(offers []string) (bestOffer string, matched bool) {
	parser := newHeaderParser(n.req.Header, false)

	return parser.selectOffer(offers, parser.parse(HeaderAcceptCharset))
}
