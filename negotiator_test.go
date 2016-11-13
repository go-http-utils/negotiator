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

	bestOffer, matched := n.Accept([]string{"TExt/htmL"})

	s.True(matched)
	s.Equal("text/html", bestOffer)
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
	s.Equal("en", bestOffer)
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
