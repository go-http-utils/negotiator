package negotiator

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParseAcceptTestSuite struct {
	suite.Suite

	header http.Header
}

func (s *ParseAcceptTestSuite) SetupTest() {
	s.header = make(http.Header)
}

func (s *ParseAcceptTestSuite) TestEmpty() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "")

	specs := parseAccept(s.header)
	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*/*", 1.0)
}

func (s *ParseAcceptTestSuite) TestAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "*/*")

	specs := parseAccept(s.header)
	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "*/*", 1.0)
}

func (s *ParseAcceptTestSuite) TestOneType() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "application/json")

	specs := parseAccept(s.header)
	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "application/json", 1.0)
}

func (s *ParseAcceptTestSuite) TestOneTypeWithQZero() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "application/json;q=0")

	specs := parseAccept(s.header)
	assert.Equal(0, len(specs))
}

func (s *ParseAcceptTestSuite) TestSortByQ() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "application/json;q=0.2, text/html")

	specs := parseAccept(s.header)
	assert.Equal(2, len(specs))

	equalSpec(assert, specs[0], "text/html", 1.0)
	equalSpec(assert, specs[1], "application/json", 0.2)
}

func (s *ParseAcceptTestSuite) TestSuffixAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "text/*")

	specs := parseAccept(s.header)
	assert.Equal(1, len(specs))

	equalSpec(assert, specs[0], "text/*", 1.0)
}

func (s *ParseAcceptTestSuite) TestSortWithAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "text/plain, application/json;q=0.5, text/html, */*;q=0.1")

	specs := parseAccept(s.header)
	assert.Equal(4, len(specs))

	equalSpec(assert, specs[0], "text/plain", 1.0)
	equalSpec(assert, specs[1], "text/html", 1.0)
	equalSpec(assert, specs[2], "application/json", 0.5)
	equalSpec(assert, specs[3], "*/*", 0.1)
}

func TestParseAccept(t *testing.T) {
	suite.Run(t, new(ParseAcceptTestSuite))
}
