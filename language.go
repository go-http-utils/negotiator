package negotiator

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func parseLanguage(header http.Header) (specs specs) {
	headerVal := strings.ToLower(strings.Replace(header.Get(HeaderAcceptLanguage),
		" ", "", -1))

	if headerVal == "" {
		specs = []spec{spec{val: "*", q: defaultQ}}
		return
	}

	languages := strings.Split(headerVal, ",")

	for _, languages := range languages {
		pair := strings.Split(strings.TrimSpace(languages), ";")

		if len(pair) < 1 || len(pair) > 2 {
			continue
		}

		spec := spec{val: pair[0], q: defaultQ}

		if len(pair) == 2 && strings.HasPrefix(pair[1], "q=") {
			var i int

			if strings.HasPrefix(pair[1], "q=") {
				i = 2
			} else if strings.HasPrefix(pair[1], "level=") {
				i = 6
			} else {
				continue
			}

			if q, err := strconv.ParseFloat(pair[1][i:], 64); err == nil && q != 0.0 {
				if q > defaultQ {
					q = defaultQ
				}

				spec.q = q
			} else {
				continue
			}
		}

		specs = append(specs, spec)
	}

	sort.Sort(specs)

	return
}
