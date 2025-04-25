package reedsolomon

import "log"

type elementOperator struct {
	a, b gfElement
}

func (op elementOperator) Add() gfElement {
	return op.a ^ op.b
}

func (op elementOperator) Sub() gfElement {
	return op.a ^ op.b
}

func (op elementOperator) Multiply() gfElement {
	if op.a == gfZero || op.b == gfZero {
		return gfZero
	}

	return gfExpTable[(gfLogTable[op.a]+gfLogTable[op.b])%255]
}

func (op elementOperator) Divide() gfElement {
	if op.a == gfZero {
		return gfZero
	} else if op.b == gfZero {
		log.Panicln("Divide by zero")
	}

	inv := op.Inverse(op.b)
	return elementOperator{a: op.a, b: inv}.Multiply()
}


func (op elementOperator) Inverse(a gfElement) gfElement {
	if a == gfZero {
		log.Panicln("No multiplicative inverse of 0")
	}

	return gfExpTable[255-gfLogTable[a]]
}
