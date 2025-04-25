package reedsolomon

import "github.com/i9si-sistemas/bitset"

func Encode(data *bitset.Bitset, numECBytes int) *bitset.Bitset {
	ecpoly := newGFPolyFromData(data)
	ecpoly = polyOperator{}.Multiply(ecpoly, newGFPolyMonomial(gfOne, numECBytes))
	generator := generatorPoly(numECBytes)
	remainder := polyOperator{}.Remainder(ecpoly, generator)
	result := bitset.Clone(data)
	result.AppendBytes(remainder.data(numECBytes))

	return result
}
