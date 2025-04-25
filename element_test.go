package reedsolomon

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestGFMultiplicationIdentities(t *testing.T) {
	for i := range 256 {
		value := gfElement(i)
		result := elementOperator{gfZero, value}.Multiply()
		assert.Equal(t, result, gfZero)
		result = elementOperator{value, gfOne}.Multiply()
		assert.Equal(t, result, value)
	}
}

func TestGFMultiplicationAndDivision(t *testing.T) {
	var tests = []struct {
		a      gfElement
		b      gfElement
		result gfElement
	}{
		{0, 29, 0},
		{1, 1, 1},
		{1, 32, 32},
		{2, 4, 8},
		{16, 128, 232},
		{17, 17, 28},
		{27, 9, 195},
	}

	for _, test := range tests {
		result := elementOperator{test.a, test.b}.Multiply()

		if result != test.result {
			t.Errorf("%d * %d = %d, want %d", test.a, test.b, result, test.result)
		}

		if test.b != gfZero && test.result != gfZero {
			b := elementOperator{test.result, test.a}.Divide()

			if b != test.b {
				t.Errorf("%d / %d = %d, want %d", test.result, test.a, b, test.b)
			}
		}
	}
}

func TestGFInverse(t *testing.T) {
	for i := 1; i < 256; i++ {
		a := gfElement(i)
		inverse := elementOperator{}.Inverse(a)

		result := elementOperator{a, inverse}.Multiply()

		if result != gfOne {
			t.Errorf("%d * %d^-1 == %d, want %d", a, inverse, result, gfOne)
		}
	}
}

func TestGFDivide(t *testing.T) {
	for i := 1; i < 256; i++ {
		for j := 1; j < 256; j++ {
			a := gfElement(i)
			b := gfElement(j)
			product := elementOperator{a, b}.Multiply()
			result := elementOperator{product, b}.Divide()

			if result != a {
				t.Errorf("%d / %d == %d, want %d", product, b, result, a)
			}
		}
	}
}
