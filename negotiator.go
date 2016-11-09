package negotiator

import "net/http"

const (
	defaultQ float64 = 1.0

	// HeaderAccept is the HTTP "Accept" Header
	HeaderAccept = "Accept"
)

// Negotiator revices a net.http.Request object
type Negotiator struct {
	req *http.Request
}

type Spec struct {
	val string
	q   float64
}
