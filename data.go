package reedsolomon

import "fmt"

func (e gfPoly) data(numTerms int) []byte {
	result := make([]byte, numTerms)

	i := numTerms - len(e.term)
	for j := len(e.term) - 1; j >= 0; j-- {
		result[i] = byte(e.term[j])
		i++
	}

	return result
}

func (e gfPoly) string(useIndexForm bool) string {
	var str string
	numTerms := e.numTerms()

	for i := numTerms - 1; i >= 0; i-- {
		if e.term[i] > 0 {
			if len(str) > 0 {
				str += " + "
			}

			if !useIndexForm {
				str += fmt.Sprintf("%dx^%d", e.term[i], i)
			} else {
				str += fmt.Sprintf("a^%dx^%d", gfLogTable[e.term[i]], i)
			}
		}
	}

	if len(str) == 0 {
		str = "0"
	}

	return str
}
