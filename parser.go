package negotiator

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type headerParser struct {
	header      http.Header
	hasSlashVal bool
	defaultQ    float64
	wildCard    string
}

func newHeaderParser(header http.Header, hasSlashVal bool) *headerParser {
	hp := &headerParser{header: header, hasSlashVal: hasSlashVal, defaultQ: 1.0}

	if hp.hasSlashVal {
		hp.wildCard = "*/*"
	} else {
		hp.wildCard = "*"
	}

	return hp
}

func (p headerParser) parse(headerName string) (specs specs) {
	headerVal := formatHeaderVal(p.header.Get(headerName))

	if headerVal == "" {
		specs = []spec{spec{val: p.wildCard, q: p.defaultQ}}
		return
	}

	for _, accept := range strings.Split(headerVal, ",") {
		pair := strings.Split(strings.TrimSpace(accept), ";")

		if len(pair) < 1 || len(pair) > 2 {
			if p.hasSlashVal {
				if strings.Index(pair[0], "/") == -1 {
					continue
				}
			} else {
				continue
			}
		}

		spec := spec{val: pair[0], q: p.defaultQ}

		if len(pair) == 2 {
			var i int

			if strings.HasPrefix(pair[1], "q=") {
				i = 2
			} else if strings.HasPrefix(pair[1], "level=") {
				i = 6
			} else {
				continue
			}

			if q, err := strconv.ParseFloat(pair[1][i:], 64); err == nil && q != 0.0 {
				if q > p.defaultQ {
					q = p.defaultQ
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

func (p headerParser) selectOffer(offers []string, specs specs) (bestOffer string) {
	bestQ := 0.0

	var bestWild, totalWild int
	if p.hasSlashVal {
		bestWild, totalWild = 3, 3
	} else {
		bestWild, totalWild = 2, 2
	}

	if len(specs) == 0 {
		return
	}

	if len(offers) == 0 {
		bestOffer = specs[0].val
		return
	}

	for _, offer := range offers {
		lowerCaseOffer := strings.ToLower(offer)

		for _, spec := range specs {
			switch {
			case spec.q <= bestQ:
				continue
			case spec.val == p.wildCard && !specs.hasVal(lowerCaseOffer):
				if spec.q > bestQ || bestWild > totalWild-1 {
					bestOffer = offer

					bestQ, bestWild = spec.q, totalWild-1
				}
			case p.hasSlashVal && strings.HasSuffix(spec.val, "/*"):
				if strings.HasPrefix(lowerCaseOffer, spec.val[:len(spec.val)-1]) &&
					(spec.q > bestQ || bestWild > totalWild-2) {
					bestOffer = offer

					bestQ, bestWild = spec.q, totalWild-2
				}
			case spec.val == lowerCaseOffer:
				if spec.q > bestQ || bestWild > 0 {
					bestOffer = offer

					bestQ, bestWild = spec.q, 0
				}
			}
		}
	}

	return
}

func formatHeaderVal(val string) string {
	return strings.ToLower(strings.Replace(val, " ", "", -1))
}
