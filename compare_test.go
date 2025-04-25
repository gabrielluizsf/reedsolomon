package reedsolomon

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestEquals(t *testing.T) {
	var tests = []struct {
		a       gfPoly
		b       gfPoly
		isEqual bool
	}{
		{
			gfPoly{[]gfElement{0}},
			gfPoly{[]gfElement{0}},
			true,
		},
		{
			gfPoly{[]gfElement{1}},
			gfPoly{[]gfElement{0}},
			false,
		},
		{
			gfPoly{[]gfElement{1, 0, 1, 0, 1}},
			gfPoly{[]gfElement{1, 0, 1, 0, 1}},
			true,
		},
		{
			gfPoly{[]gfElement{1, 0, 1}},
			gfPoly{[]gfElement{1, 0, 1, 0, 0}},
			true,
		},
	}

	for _, test := range tests {
		isEqual := test.a.equals(test.b)
		assert.Equal(t, isEqual, test.isEqual)
	}
}
