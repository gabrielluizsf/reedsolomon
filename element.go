package reedsolomon

type gfElement uint8

const (
	gfZero = gfElement(0)
	gfOne  = gfElement(1)
)

func Element(data byte) gfElement {
	return gfElement(data)
}