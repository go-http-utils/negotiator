package negotiator

import (
	"github.com/stretchr/testify/assert"
)

func equalSpec(assert *assert.Assertions, spec spec, val string, q float64) {
	assert.Equal(val, spec.val)
	assert.Equal(q, spec.q)
}
