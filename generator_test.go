package reedsolomon

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestGeneratorPoly(t *testing.T) {
	var tests = []struct {
		degree    int
		generator gfPoly
	}{
		{
			2,
			gfPoly{term: []gfElement{2, 3, 1}},
		},
		{
			5,
			gfPoly{term: []gfElement{116, 147, 63, 198, 31, 1}},
		},
		{
			68,
			gfPoly{term: []gfElement{11, 99, 29, 32, 8, 204, 149, 34, 12,
				235, 11, 119, 7, 255, 239, 211, 157, 80, 4, 199, 36, 63, 88, 158, 51, 212,
				219, 20, 245, 226, 175, 14, 20, 144, 225, 230, 246, 71, 107, 38, 107, 182,
				170, 224, 172, 145, 112, 185, 20, 109, 167, 174, 34, 107, 85, 119, 142,
				157, 230, 223, 94, 60, 182, 18, 39, 9, 115, 131, 1}},
		},
	}

	for _, test := range tests {
		generator := generatorPoly(test.degree)
		assert.True(t, generator.equals(test.generator))
	}
}
