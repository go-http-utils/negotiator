package negotiator

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func parseAccept(header http.Header) (specs specs) {
	headerVal := header.Get(HeaderAccept)

	if headerVal == "" {
		specs = []spec{spec{val: "*/*", q: defaultQ}}
		return
	}

	accpets := strings.Split(headerVal, ",")

	for _, accept := range accpets {
		pair := strings.Split(strings.TrimSpace(accept), ";")

		if len(pair) < 1 || len(pair) > 2 {
			continue
		}

		val := strings.TrimSpace(pair[0])
		i := strings.Index(val, "/")

		if i == -1 {
			continue
		}

		spec := spec{val: val, q: defaultQ}

		if len(pair) == 2 && strings.HasPrefix(pair[1], "q=") {
			var i int

			if strings.HasPrefix(pair[1], "q=") {
				i = 2
			} else if strings.HasPrefix(pair[1], "level=") {
				i = 6
			} else {
				continue
			}

			if q, err := strconv.ParseFloat(pair[1][i:], 64); err == nil && q != 0 {
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
