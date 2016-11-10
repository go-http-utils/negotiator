package negotiator

import (
	"net/http"
	"strings"
)

const (
	defaultQ float64 = 1.0

	// HeaderAccept is the HTTP "Accept" Header
	HeaderAccept = "Accept"
)

// Negotiator revices a net.http.Request object
type Negotiator struct {
	req *http.Request
}

type spec struct {
	val string
	q   float64
}

// Specs it the shorthand for []Spec
type specs []spec

// Len is used to impelement sort.Interface for Specs
func (ss specs) Len() int {
	return len(ss)
}

// Swap is used to impelement sort.Interface for Specs
func (ss specs) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

// Less is used to impelement sort.Interface for Specs
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
