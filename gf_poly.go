package reedsolomon

import "github.com/i9si-sistemas/bitset"

func newGFPolyFromData(data *bitset.Bitset) gfPoly {
	numTotalBytes := data.Len() / 8
	if data.Len()%8 != 0 {
		numTotalBytes++
	}

	result := gfPoly{term: make([]gfElement, numTotalBytes)}

	i := numTotalBytes - 1
	for j := 0; j < data.Len(); j += 8 {
		result.term[i] = gfElement(data.ByteAt(j))
		i--
	}

	return result
}

func newGFPolyMonomial(term gfElement, degree int) gfPoly {
	if term == gfZero {
		return gfPoly{}
	}

	result := gfPoly{term: make([]gfElement, degree+1)}
	result.term[degree] = term

	return result
}
