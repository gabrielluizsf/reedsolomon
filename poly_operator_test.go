package reedsolomon

import "testing"

var operator = NewGFPolyOperator()

func TestGFPolyAdd(t *testing.T) {
	var tests = []struct {
		a      gfPoly
		b      gfPoly
		result gfPoly
	}{
		{
			gfPoly{[]gfElement{0, 0, 0}},
			gfPoly{[]gfElement{0}},
			gfPoly{[]gfElement{}},
		},
		{
			gfPoly{[]gfElement{1, 0}},
			gfPoly{[]gfElement{1, 0}},
			gfPoly{[]gfElement{0, 0}},
		},
		{
			gfPoly{[]gfElement{0xA0, 0x80, 0xFF, 0x00}},
			gfPoly{[]gfElement{0x0A, 0x82}},
			gfPoly{[]gfElement{0xAA, 0x02, 0xFF}},
		},
	}

	for _, test := range tests {
		result := operator.Add(test.a, test.b)

		if !test.result.equals(result) {
			t.Errorf("%s * %s != %s (got %s)\n", test.a.string(false), test.b.string(false),
				test.result.string(false), result.string(false))
		}

		if len(result.term) > 0 && result.term[len(result.term)-1] == 0 {
			t.Errorf("Result's maximum term coefficient is zero")
		}
	}
}

func TestGFPolyMultiply(t *testing.T) {
	var tests = []struct {
		a      gfPoly
		b      gfPoly
		result gfPoly
	}{
		{
			gfPoly{[]gfElement{0, 0, 1}},
			gfPoly{[]gfElement{9}},
			gfPoly{[]gfElement{0, 0, 9}},
		},
		{
			gfPoly{[]gfElement{0, 16, 1}},
			gfPoly{[]gfElement{128, 2}},
			gfPoly{[]gfElement{0, 232, 160, 2}},
		},
		{
			gfPoly{[]gfElement{254, 120, 88, 44, 11, 1}},
			gfPoly{[]gfElement{16, 2, 0, 51, 44}},
			gfPoly{[]gfElement{91, 50, 25, 184, 194, 105, 45, 244, 58, 44}},
		},
	}

	for _, test := range tests {
		result := operator.Multiply(test.a, test.b)

		if !test.result.equals(result) {
			t.Errorf("%s * %s = %s (got %s)\n",
				test.a.string(false),
				test.b.string(false),
				test.result.string(false),
				result.string(false))
		}
	}
}

func TestGFPolyRemainder(t *testing.T) {
	var tests = []struct {
		numerator   gfPoly
		denominator gfPoly
		remainder   gfPoly
	}{
		{
			gfPoly{[]gfElement{1}},
			gfPoly{[]gfElement{1}},
			gfPoly{[]gfElement{0}},
		},
		{
			gfPoly{[]gfElement{1, 0}},
			gfPoly{[]gfElement{1}},
			gfPoly{[]gfElement{0}},
		},
		{
			gfPoly{[]gfElement{1}},
			gfPoly{[]gfElement{1, 0}},
			gfPoly{[]gfElement{1}},
		},
		{
			gfPoly{[]gfElement{1, 0, 1}},
			gfPoly{[]gfElement{0, 1}},
			gfPoly{[]gfElement{1}},
		},
		{
			gfPoly{[]gfElement{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}},
			gfPoly{[]gfElement{1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1}},
			gfPoly{[]gfElement{0, 0, 1, 1, 1, 0, 1, 1}},
		},
		{
			gfPoly{[]gfElement{91, 50, 25, 184, 194, 105, 45, 244, 58, 44}},
			gfPoly{[]gfElement{254, 120, 88, 44, 11, 1}},
			gfPoly{[]gfElement{}},
		},
		{
			gfPoly{[]gfElement{0, 0, 0, 0, 0, 0, 195, 172, 24, 64}},
			gfPoly{[]gfElement{116, 147, 63, 198, 31, 1}},
			gfPoly{[]gfElement{48, 174, 34, 13, 134}},
		},
	}

	for _, test := range tests {
		remainder := operator.Remainder(test.numerator, test.denominator)

		if !test.remainder.equals(remainder) {
			t.Errorf("%s / %s, remainder = %s (got %s)\n",
				test.numerator.string(false),
				test.denominator.string(false),
				test.remainder.string(false),
				remainder.string(false))
		}
	}
}
