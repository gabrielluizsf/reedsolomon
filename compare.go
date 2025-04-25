package reedsolomon

func (e gfPoly) equals(other gfPoly) bool {
	var minecPoly *gfPoly
	var maxecPoly *gfPoly

	if e.numTerms() > other.numTerms() {
		minecPoly = &other
		maxecPoly = &e
	} else {
		minecPoly = &e
		maxecPoly = &other
	}

	numMinTerms := minecPoly.numTerms()
	numMaxTerms := maxecPoly.numTerms()

	for i := range numMinTerms {
		if e.term[i] != other.term[i] {
			return false
		}
	}

	for i := numMinTerms; i < numMaxTerms; i++ {
		if maxecPoly.term[i] != 0 {
			return false
		}
	}

	return true
}
