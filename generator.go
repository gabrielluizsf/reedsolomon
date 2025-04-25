package reedsolomon

import "log"

type gfPoly struct {
	term []gfElement
}

func generatorPoly(degree int) gfPoly {
	if degree < 2 {
		log.Panic("degree < 2")
	}

	gen := gfPoly{term: []gfElement{1}}

	for i := range degree {
		nextPoly := gfPoly{term: []gfElement{gfExpTable[i], 1}}
		gen = polyOperator{}.Multiply(gen, nextPoly)
	}

	return gen
}
