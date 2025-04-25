package reedsolomon

import "log"

type polyOperator struct {}

func NewGFPolyOperator() polyOperator {
	return polyOperator{}
}

func (op polyOperator) Multiply(a, b gfPoly) gfPoly {
	numATerms := a.numTerms()
	numBTerms := b.numTerms()

	result := gfPoly{term: make([]gfElement, numATerms+numBTerms)}

	for i := range numATerms {
		for j := range numBTerms {
			if a.term[i] != 0 && b.term[j] != 0 {
				monomial := gfPoly{term: make([]gfElement, i+j+1)}
				monomial.term[i+j] = elementOperator{a.term[i], b.term[j]}.Multiply()

				result = op.Add(result, monomial)
			}
		}
	}

	return result.normalised()
}

func (op polyOperator) Remainder(numerator, denominator gfPoly) gfPoly {
	if denominator.equals(gfPoly{}) {
		log.Panicln("Remainder by zero")
	}

	remainder := numerator

	for remainder.numTerms() >= denominator.numTerms() {
		degree := remainder.numTerms() - denominator.numTerms()
		coefficient := elementOperator{
					remainder.term[remainder.numTerms()-1],
					denominator.term[denominator.numTerms()-1],
		}.Divide()

		divisor := op.Multiply(denominator,
			newGFPolyMonomial(coefficient, degree))

		remainder = op.Add(remainder, divisor)
	}

	return remainder.normalised()
}

func (op polyOperator) Add(a, b gfPoly) gfPoly {
	numATerms := a.numTerms()
	numBTerms := b.numTerms()

	numTerms := numATerms
	if numBTerms > numTerms {
		numTerms = numBTerms
	}

	result := gfPoly{term: make([]gfElement, numTerms)}

	for i := range  numTerms {
		switch {
		case numATerms > i && numBTerms > i:
			result.term[i] = elementOperator{a.term[i], b.term[i]}.Add()
		case numATerms > i:
			result.term[i] = a.term[i]
		default:
			result.term[i] = b.term[i]
		}
	}

	return result.normalised()
}
