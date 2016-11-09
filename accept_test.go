package negotiator

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AcceptTestSuite struct {
	suite.Suite

	header http.Header
}

func (s *AcceptTestSuite) SetupTest() {
	s.header = make(http.Header)
}

func (s *AcceptTestSuite) TestEmpty() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "")

	specs := ParseAccept(s.header)
	assert.Equal(1, len(specs))
	assert.Equal("*/*", specs[0].val)
	assert.Equal(1.0, specs[0].q)
}

func (s *AcceptTestSuite) TestAsterisk() {
	assert := assert.New(s.T())

	s.header.Set(HeaderAccept, "*/*")

	specs := ParseAccept(s.header)
	assert.Equal(1, len(specs))
	assert.Equal("*/*", specs[0].val)
	assert.Equal(1.0, specs[0].q)
}

func TestAccept(t *testing.T) {
	suite.Run(t, new(AcceptTestSuite))
}
