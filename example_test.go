package negotiator_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-http-utils/negotiator"
)

func ExampleNegotiator_Accept() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "text/html, application/*;q=0.9, image/jpeg;q=0.8")
	negotiator := negotiator.New(req.Header)

	fmt.Println(negotiator.Type("text/html", "application/json", "image/jpeg"))
	// -> "text/html"

	fmt.Println(negotiator.Type("application/json", "image/jpeg", "text/plain"))
	// -> "application/json"

	fmt.Println(negotiator.Type("text/plain"))
	// -> ""
}

func ExampleNegotiator_Encoding() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Encoding", "gzip, compress;q=0.2, identity;q=0.5")
	negotiator := negotiator.New(req.Header)

	fmt.Println(negotiator.Encoding("identity", "gzip"))
	// -> "gzip"

	fmt.Println(negotiator.Encoding("compress", "identity"))
	// -> "identity"
}

func ExampleNegotiator_Language() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "en;q=0.8, es, pt")
	negotiator := negotiator.New(req.Header)

	fmt.Println(negotiator.Language("en", "es", "fr"))
	// -> "es"

	fmt.Println(negotiator.Language("es", "pt"))
	// -> "es"
}

func ExampleNegotiator_Charset() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "utf-8, iso-8859-1;q=0.8, utf-7;q=0.2")
	negotiator := negotiator.New(req.Header)

	fmt.Println(negotiator.Charset("UTF-8", "ISO-8859-1", "ISO-8859-5"))
	// -> "UTF-8"

	fmt.Println(negotiator.Charset("ISO-8859-5"))
	// -> ""
}
