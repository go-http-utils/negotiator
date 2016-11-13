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

func setUpNegotiator(header, val string) Negotiator {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(header, val)

	return New(req)
}

// Accept
type AcceptSuite struct {
	suite.Suite
}

func (s AcceptSuite) TestEmpty() {
	n := setUpNegotiator(HeaderAccept, "application/json;q=0.2, text/html")

	_, matched := n.Accept([]string{})

	s.False(matched)
}

func (s AcceptSuite) TestCaseInsensitive() {
	n := setUpNegotiator(HeaderAccept, "text/html")

	bestOffer, matched := n.Accept([]string{"TEXT/HTML"})

	s.True(matched)
	s.Equal("TEXT/HTML", bestOffer)
}

func (s AcceptSuite) TestUnMatched() {
	n := setUpNegotiator(HeaderAccept, "application/json;q=0.2, text/html")

	_, matched := n.Accept([]string{"text/plain"})

	s.False(matched)
}

func (s AcceptSuite) TestEmptyAccepts() {
	n := setUpNegotiator(HeaderAccept, "application/json;q=0")

	_, matched := n.Accept([]string{"application/json"})

	s.False(matched)
}

func (s AcceptSuite) TestOneMatch() {
	n := setUpNegotiator(HeaderAccept, "application/json;q=0.2")

	bestOffer, matched := n.Accept([]string{"application/json"})

	s.True(matched)
	s.Equal("application/json", bestOffer)
}

func (s AcceptSuite) TestWithAsterisk() {
	n := setUpNegotiator(HeaderAccept, "text/*")

	bestOffer, matched := n.Accept([]string{"text/*"})

	s.True(matched)
	s.Equal("text/*", bestOffer)
}

func (s AcceptSuite) TestMatchAsterisk() {
	n := setUpNegotiator(HeaderAccept, "text/*")

	bestOffer, matched := n.Accept([]string{"text/html"})

	s.True(matched)
	s.Equal("text/html", bestOffer)
}

func (s AcceptSuite) TestFirstMatchAsterisk() {
	n := setUpNegotiator(HeaderAccept, "text/*")

	bestOffer, matched := n.Accept([]string{"text/html", "text/plain", "application/json"})

	s.True(matched)
	s.Equal("text/html", bestOffer)
}

func (s AcceptSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(HeaderAccept, "*/*, application/json;q=0.2")

	bestOffer, matched := n.Accept([]string{"text/html", "application/json", "text/plain"})

	s.True(matched)
	s.Equal("text/html", bestOffer)
}

func (s AcceptSuite) TestWithAllAsterisk() {
	n := setUpNegotiator(HeaderAccept, "*/*")

	bestOffer, matched := n.Accept([]string{"application/json", "text/html", "text/plain"})

	s.True(matched)
	s.Equal("application/json", bestOffer)
}

func TestAccept(t *testing.T) {
	suite.Run(t, new(AcceptSuite))
}

// Language
type LanguageSuite struct {
	suite.Suite
}

func (s LanguageSuite) TestEmpty() {
	n := setUpNegotiator(HeaderAcceptLanguage, "")

	_, matched := n.Language([]string{})

	s.False(matched)
}

func (s LanguageSuite) TestCaseInsensitive() {
	n := setUpNegotiator(HeaderAcceptLanguage, "En")

	bestOffer, matched := n.Language([]string{"eN"})

	s.True(matched)
	s.Equal("eN", bestOffer)
}

func (s LanguageSuite) TestUnMatched() {
	n := setUpNegotiator(HeaderAcceptLanguage, "en,zh")

	_, matched := n.Language([]string{"ko"})

	s.False(matched)
}

func (s LanguageSuite) TestEmptyLanguages() {
	n := setUpNegotiator(HeaderAcceptLanguage, "en;q=0")

	_, matched := n.Language([]string{"en"})

	s.False(matched)
}

func (s LanguageSuite) TestOneMatch() {
	n := setUpNegotiator(HeaderAcceptLanguage, "en;q=0.2")

	bestOffer, matched := n.Language([]string{"en"})

	s.True(matched)
	s.Equal("en", bestOffer)
}

func (s LanguageSuite) TestMatchAsterisk() {
	n := setUpNegotiator(HeaderAcceptLanguage, "*")

	bestOffer, matched := n.Language([]string{"ko", "en"})

	s.True(matched)
	s.Equal("ko", bestOffer)
}

func (s LanguageSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(HeaderAcceptLanguage, "*, ko;q=0.5")

	bestOffer, matched := n.Language([]string{"en", "ko", "zh"})

	s.True(matched)
	s.Equal("en", bestOffer)
}

func TestLanguage(t *testing.T) {
	suite.Run(t, new(LanguageSuite))
}

// Encoding
type EncodingSuite struct {
	suite.Suite
}

func (s EncodingSuite) TestEmpty() {
	n := setUpNegotiator(HeaderAcceptEncoding, "")

	_, matched := n.Encoding([]string{})

	s.False(matched)
}

func (s EncodingSuite) TestCaseInsensitive() {
	n := setUpNegotiator(HeaderAcceptEncoding, "GZip")

	bestOffer, matched := n.Encoding([]string{"Gzip"})

	s.True(matched)
	s.Equal("Gzip", bestOffer)
}

func (s EncodingSuite) TestUnMatched() {
	n := setUpNegotiator(HeaderAcceptEncoding, "gzip,default")

	_, matched := n.Encoding([]string{"zlib"})

	s.False(matched)
}

func (s EncodingSuite) TestEmptyLanguages() {
	n := setUpNegotiator(HeaderAcceptEncoding, "gzip;q=0")

	_, matched := n.Encoding([]string{"gzip"})

	s.False(matched)
}

func (s EncodingSuite) TestOneMatch() {
	n := setUpNegotiator(HeaderAcceptEncoding, "gzip;q=0.2")

	bestOffer, matched := n.Encoding([]string{"gzip"})

	s.True(matched)
	s.Equal("gzip", bestOffer)
}

func (s EncodingSuite) TestMatchAsterisk() {
	n := setUpNegotiator(HeaderAcceptEncoding, "*")

	bestOffer, matched := n.Encoding([]string{"gzip", "deflate"})

	s.True(matched)
	s.Equal("gzip", bestOffer)
}

func (s EncodingSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(HeaderAcceptEncoding, "*, gzip;q=0.5")

	bestOffer, matched := n.Encoding([]string{"gzip", "deflate", "zlib"})

	s.True(matched)
	s.Equal("deflate", bestOffer)
}

func TestEncoding(t *testing.T) {
	suite.Run(t, new(EncodingSuite))
}

// Charset
type CharsetSuite struct {
	suite.Suite
}

func (s CharsetSuite) TestEmpty() {
	n := setUpNegotiator(HeaderAcceptCharset, "")

	_, matched := n.Charset([]string{})

	s.False(matched)
}

func (s CharsetSuite) TestCaseInsensitive() {
	n := setUpNegotiator(HeaderAcceptCharset, "ISO-8859-1")

	bestOffer, matched := n.Charset([]string{"ISO-8859-1"})

	s.True(matched)
	s.Equal("ISO-8859-1", bestOffer)
}

func (s CharsetSuite) TestUnMatched() {
	n := setUpNegotiator(HeaderAcceptCharset, "ISO-8859-1,UTF-8")

	_, matched := n.Charset([]string{"ASCII"})

	s.False(matched)
}

func (s CharsetSuite) TestEmptyCharset() {
	n := setUpNegotiator(HeaderAcceptCharset, "UTF-8;q=0")

	_, matched := n.Charset([]string{"UTF-8"})

	s.False(matched)
}

func (s CharsetSuite) TestOneMatch() {
	n := setUpNegotiator(HeaderAcceptCharset, "UTF-8;q=0.2")

	bestOffer, matched := n.Charset([]string{"UTF-8"})

	s.True(matched)
	s.Equal("UTF-8", bestOffer)
}

func (s CharsetSuite) TestMatchAsterisk() {
	n := setUpNegotiator(HeaderAcceptCharset, "*")

	bestOffer, matched := n.Charset([]string{"UTF-8", "ISO-8859-1"})

	s.True(matched)
	s.Equal("UTF-8", bestOffer)
}

func (s CharsetSuite) TestFirstMatchAllAsterisk() {
	n := setUpNegotiator(HeaderAcceptCharset, "*, UTF-8;q=0.5")

	bestOffer, matched := n.Charset([]string{"UTF-8", "ISO-8859-1", "ASCII"})

	s.True(matched)
	s.Equal("ISO-8859-1", bestOffer)
}

func (s CharsetSuite) TestHighOrderPreferred() {
	n := setUpNegotiator(HeaderAcceptCharset, "UTF-8;q=0.6, ISO-8859-1;q=0.8, UTF-8;q=0.9")

	bestOffer, matched := n.Charset([]string{"UTF-8", "ISO-8859-1", "ASCII"})

	s.True(matched)
	s.Equal("UTF-8", bestOffer)
}

func TestCharset(t *testing.T) {
	suite.Run(t, new(CharsetSuite))
}
