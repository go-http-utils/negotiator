package negotiator

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseEncodingTestSuite struct {
	suite.Suite

	parser *headerParser
	header http.Header
}

func (s *ParseEncodingTestSuite) SetupTest() {
	s.header = make(http.Header)
	s.parser = newHeaderParser(s.header, false)
}

func (s *ParseEncodingTestSuite) TestEmpty() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptEncoding, "")
	specs := s.parser.parse(headerAcceptEncoding)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseEncodingTestSuite) TestAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptEncoding, "*")
	specs := s.parser.parse(headerAcceptEncoding)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseEncodingTestSuite) TestOneEncoing() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptEncoding, "gzip")
	specs := s.parser.parse(headerAcceptEncoding)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "gzip", 1.0)
}

func (s *ParseEncodingTestSuite) TestOneEncodingWithQZero() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptEncoding, "*, gzip;q=0")
	specs := s.parser.parse(headerAcceptEncoding)

	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*", 1.0)
}

func (s *ParseEncodingTestSuite) TestSortByQ() {
	assert := assert.New(s.T())

	s.header.Set(headerAcceptEncoding, "*;q=0.8, defalte, gzip")
	specs := s.parser.parse(headerAcceptEncoding)

	assert.Equal(3, len(specs))

	equalSpec(assert, specs[0], "defalte", 1.0)
	equalSpec(assert, specs[1], "gzip", 1.0)
	equalSpec(assert, specs[2], "*", 0.8)
}

func TestParseEncoding(t *testing.T) {
	suite.Run(t, new(ParseEncodingTestSuite))
}
