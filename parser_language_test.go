package negotiator

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseLanguageTestSuite struct {
	suite.Suite

	parser *headerParser
	header http.Header
}

func (s *ParseLanguageTestSuite) SetupTest() {
	s.header = make(http.Header)
	s.parser = newHeaderParser(s.header, false)
}

func (s *ParseLanguageTestSuite) TestEmpty() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptLanguage, "")
	specs := s.parser.parse(headerAcceptLanguage)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseLanguageTestSuite) TestAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptLanguage, "*")
	specs := s.parser.parse(headerAcceptLanguage)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseLanguageTestSuite) TestOneLanguage() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptLanguage, "en")
	specs := s.parser.parse(headerAcceptLanguage)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "en", 1.0)
}

func (s *ParseLanguageTestSuite) TestOneLanguageWithQZero() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptLanguage, "*, en;q=0")
	specs := s.parser.parse(headerAcceptLanguage)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseLanguageTestSuite) TestSortByQ() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptLanguage, "*;q=0.8, en, es")
	specs := s.parser.parse(headerAcceptLanguage)

	assert.Equal(3, len(specs))

	equalSpec(assert, specs[0], "en", 1.0)
	equalSpec(assert, specs[1], "es", 1.0)
	equalSpec(assert, specs[2], "*", 0.8)
}

func TestParseLanguage(t *testing.T) {
	suite.Run(t, new(ParseLanguageTestSuite))
}
