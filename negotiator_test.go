package negotiator

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func equalSpec(assert *assert.Assertions, spec spec, val string, q float64) {
	assert.Equal(val, spec.val)
	assert.Equal(q, spec.q)
}

func setUpNegotiator(header, val string) *Negotiator {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(header, val)

	return New(req.Header)
}

// Accept
type AcceptSuite struct {
	suite.Suite
}

func (s AcceptSuite) TestEmpty() {
	n := setUpNegotiator(headerAccept, "application/json;q=0.2, text/html")
	s.Equal("text/html", n.Type())
}

func (s AcceptSuite) TestCaseInsensitive() {
	n := setUpNegotiator(headerAccept, "text/html")
	s.Equal("TEXT/HTML", n.Type("TEXT/HTML"))
}

func (s AcceptSuite) TestUnMatched() {
	n := setUpNegotiator(headerAccept, "application/json;q=0.2, text/html")
	s.Equal("", n.Type("text/plain"))
}

func (s AcceptSuite) TestEmptyAccepts() {
	n := setUpNegotiator(headerAccept, "application/json;q=0")
	s.Equal("", n.Type("application/json"))
}

func (s AcceptSuite) TestOneMatch() {
	n := setUpNegotiator(headerAccept, "application/json;q=0.2")
	s.Equal("application/json", n.Type("application/json"))
}

func (s AcceptSuite) TestWithAsterisk() {
	n := setUpNegotiator(headerAccept, "text/*")
	s.Equal("text/*", n.Type("text/*"))
}

func (s AcceptSuite) TestMatchAsterisk() {
	n := setUpNegotiator(headerAccept, "text/*")
	s.Equal("text/html", n.Type("text/html"))
}

func (s AcceptSuite) TestFirstMatchAsterisk() {
	n := setUpNegotiator(headerAccept, "text/*")
	s.Equal("text/html", n.Type("text/html", "text/plain", "application/json"))
}

func (s AcceptSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(headerAccept, "*/*, application/json;q=0.2")
	s.Equal("text/html", n.Type("text/html", "application/json", "text/plain"))
}

func (s AcceptSuite) TestWithAllAsterisk() {
	n := setUpNegotiator(headerAccept, "*/*")
	s.Equal("application/json", n.Type("application/json", "text/html", "text/plain"))
}

func TestAccept(t *testing.T) {
	suite.Run(t, new(AcceptSuite))
}

// Language
type LanguageSuite struct {
	suite.Suite
}

func (s LanguageSuite) TestEmpty() {
	n := setUpNegotiator(headerAcceptLanguage, "")
	s.Equal("*", n.Language())
}

func (s LanguageSuite) TestCaseInsensitive() {
	n := setUpNegotiator(headerAcceptLanguage, "En")
	s.Equal("eN", n.Language("eN"))
}

func (s LanguageSuite) TestUnMatched() {
	n := setUpNegotiator(headerAcceptLanguage, "en,zh")
	s.Equal("", n.Language("ko"))
}

func (s LanguageSuite) TestEmptyLanguages() {
	n := setUpNegotiator(headerAcceptLanguage, "en;q=0")
	s.Equal("", n.Language("en"))
}

func (s LanguageSuite) TestOneMatch() {
	n := setUpNegotiator(headerAcceptLanguage, "en;q=0.2")
	s.Equal("en", n.Language("en"))
}

func (s LanguageSuite) TestMatchAsterisk() {
	n := setUpNegotiator(headerAcceptLanguage, "*")
	s.Equal("ko", n.Language("ko", "en"))
}

func (s LanguageSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(headerAcceptLanguage, "*, ko;q=0.5")
	s.Equal("en", n.Language("en", "ko", "zh"))
}

func TestLanguage(t *testing.T) {
	suite.Run(t, new(LanguageSuite))
}

// Encoding
type EncodingSuite struct {
	suite.Suite
}

func (s EncodingSuite) TestEmpty() {
	n := setUpNegotiator(headerAcceptEncoding, "")
	s.Equal("*", n.Encoding())
}

func (s EncodingSuite) TestCaseInsensitive() {
	n := setUpNegotiator(headerAcceptEncoding, "GZip")
	s.Equal("Gzip", n.Encoding("Gzip"))
}

func (s EncodingSuite) TestUnMatched() {
	n := setUpNegotiator(headerAcceptEncoding, "gzip,default")
	s.Equal("", n.Encoding("zlib"))
}

func (s EncodingSuite) TestEmptyLanguages() {
	n := setUpNegotiator(headerAcceptEncoding, "gzip;q=0")
	s.Equal("", n.Encoding("gzip"))
}

func (s EncodingSuite) TestOneMatch() {
	n := setUpNegotiator(headerAcceptEncoding, "gzip;q=0.2")
	s.Equal("gzip", n.Encoding("gzip"))
}

func (s EncodingSuite) TestMatchAsterisk() {
	n := setUpNegotiator(headerAcceptEncoding, "*")
	s.Equal("gzip", n.Encoding("gzip", "deflate"))
}

func (s EncodingSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(headerAcceptEncoding, "*, gzip;q=0.5")
	s.Equal("deflate", n.Encoding("gzip", "deflate", "zlib"))
}

func TestEncoding(t *testing.T) {
	suite.Run(t, new(EncodingSuite))
}

// Charset
type CharsetSuite struct {
	suite.Suite
}

func (s CharsetSuite) TestEmpty() {
	n := setUpNegotiator(headerAcceptCharset, "")
	s.Equal("*", n.Charset())
}

func (s CharsetSuite) TestCaseInsensitive() {
	n := setUpNegotiator(headerAcceptCharset, "ISO-8859-1")
	s.Equal("ISO-8859-1", n.Charset("ISO-8859-1"))
}

func (s CharsetSuite) TestUnMatched() {
	n := setUpNegotiator(headerAcceptCharset, "ISO-8859-1,UTF-8")
	s.Equal("", n.Charset("ASCII"))
}

func (s CharsetSuite) TestEmptyCharset() {
	n := setUpNegotiator(headerAcceptCharset, "UTF-8;q=0")
	s.Equal("", n.Charset("UTF-8"))
}

func (s CharsetSuite) TestOneMatch() {
	n := setUpNegotiator(headerAcceptCharset, "UTF-8;q=0.2")
	s.Equal("UTF-8", n.Charset("UTF-8"))
}

func (s CharsetSuite) TestMatchAsterisk() {
	n := setUpNegotiator(headerAcceptCharset, "*")
	s.Equal("UTF-8", n.Charset("UTF-8", "ISO-8859-1"))
}

func (s CharsetSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(headerAcceptCharset, "*, UTF-8;q=0.5")
	s.Equal("ISO-8859-1", n.Charset("UTF-8", "ISO-8859-1", "ASCII"))
}

func (s CharsetSuite) TestHighOrderPreferred() {
	n := setUpNegotiator(headerAcceptCharset, "UTF-8;q=0.6, ISO-8859-1;q=0.8, UTF-8;q=0.9")
	s.Equal("UTF-8", n.Charset("UTF-8", "ISO-8859-1", "ASCII"))
}

func TestCharset(t *testing.T) {
	suite.Run(t, new(CharsetSuite))
}
