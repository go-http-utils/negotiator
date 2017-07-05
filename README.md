# negotiator

[![Build Status](https://travis-ci.org/go-http-utils/negotiator.svg?branch=master)](https://travis-ci.org/go-http-utils/negotiator)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/negotiator/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/negotiator?branch=master)

An HTTP content negotiator for Go

## Installation

```sh
go get -u github.com/go-http-utils/negotiator
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/go-http-utils/negotiator

## Usage

```go
import (
  "github.com/go-http-utils/negotiator"
)

negotiator := negotiator.New(req.Header)
```

### Type

```go
// Assume that the Accept header is "text/html, application/*;q=0.9, image/jpeg;q=0.8"

negotiator.Type()
// -> "text/html"

negotiator.Type("text/html", "application/json", "image/jpeg")
// -> "text/html"

negotiator.Type("application/json", "image/jpeg", "text/plain")
// -> "application/json"

negotiator.Type("text/plain")
// -> ""
```

### Encoding

```go
// Assume that the Accept-Encoding header is "gzip, compress;q=0.2, identity;q=0.5"

negotiator.Encoding()
// -> "gzip"

negotiator.Encoding("identity", "gzip")
// -> "gzip"

negotiator.Encoding("compress", "identity")
// -> "identity"
```

### Language

```go
// Assume that the Accept-Language header is "en;q=0.8, es, pt"

negotiator.Language()
// -> "es"

negotiator.Language("en", "es", "fr")
// -> "es"

negotiator.Language("es", "pt")
// -> "es"
```

### Charset

```go
// Assume that the Accept-Charset header is "utf-8, iso-8859-1;q=0.8, utf-7;q=0.2"

negotiator.Charset()
// -> "utf-8"

negotiator.Charset("utf-8", "iso-8859-1", "iso-8859-5")
// -> "utf-8"

negotiator.Charset("iso-8859-5")
// -> ""
```
