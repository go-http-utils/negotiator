package negotiator

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseCharsetTestSuite struct {
	suite.Suite

	parser *headerParser
	header http.Header
}

func (s *ParseCharsetTestSuite) SetupTest() {
	s.header = make(http.Header)
	s.parser = newHeaderParser(s.header, false)
}

func (s *ParseCharsetTestSuite) TestEmpty() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptCharset, "")
	specs := s.parser.parse(headerAcceptCharset)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseCharsetTestSuite) TestAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptCharset, "*")
	specs := s.parser.parse(headerAcceptCharset)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseCharsetTestSuite) TestOneLanguage() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptCharset, "UTF-8;level=1.0")
	specs := s.parser.parse(headerAcceptCharset)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "utf-8", 1.0)
}

func (s *ParseCharsetTestSuite) TestOneLanguageWithQZero() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptCharset, "*, ISO-8859-1;level=0")
	specs := s.parser.parse(headerAcceptCharset)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseCharsetTestSuite) TestSortByQ() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptCharset, "*;level=0.8, ISO-8859-1, UTF-8")
	specs := s.parser.parse(headerAcceptCharset)

	assert.Equal(3, len(specs))

	equalSpec(assert, specs[0], "iso-8859-1", 1.0)
	equalSpec(assert, specs[1], "utf-8", 1.0)
	equalSpec(assert, specs[2], "*", 0.8)
}

func TestParseCharset(t *testing.T) {
	suite.Run(t, new(ParseCharsetTestSuite))
}
