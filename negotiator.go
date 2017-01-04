package negotiator

import (
	"net/http"
	"strings"

	"github.com/go-http-utils/headers"
)

// Version is this package's version
const Version = "0.1.0"

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
	header http.Header
}

// New creates an instance of Negotiator.
func New(header http.Header) *Negotiator {
	return &Negotiator{header}
}

// Type returns the most preferred content type from the HTTP Accept header.
// If nothing accepted, then empty string is returned.
func (n *Negotiator) Type(offers ...string) (bestOffer string) {
	parser := newHeaderParser(n.header, true)
	return parser.selectOffer(offers, parser.parse(headers.Accept))
}

// Language returns the most preferred language from the HTTP Accept-Language
// header. If nothing accepted, then empty string is returned.
func (n *Negotiator) Language(offers ...string) (bestOffer string) {
	parser := newHeaderParser(n.header, false)
	return parser.selectOffer(offers, parser.parse(headers.AcceptLanguage))
}

// Encoding returns the most preferred language from the HTTP Accept-Encoding
// header. If nothing accepted, then empty string is returned.
func (n *Negotiator) Encoding(offers ...string) (bestOffer string) {
	parser := newHeaderParser(n.header, false)
	return parser.selectOffer(offers, parser.parse(headers.AcceptEncoding))
}

// Charset returns the most preferred language from the HTTP Accept-Charset
// header. If nothing accepted, then empty string is returned.
func (n *Negotiator) Charset(offers ...string) (bestOffer string) {
	parser := newHeaderParser(n.header, false)
	return parser.selectOffer(offers, parser.parse(headers.AcceptCharset))
}
