package main

import "fmt"

func main() {
	var a interface{}

	a = 10
	// %T is TYPE and %v is VALUE
	fmt.Printf("%T %v\n", a, a)

	a = "Hello"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string {
		return "Hello"
	}
	fmt.Printf("%T %v\n", a, a)

	c := cube{edge: 10}
	c.Volume()
	v := VolumeOf(c)
	fmt.Printf("%T %v\n", v, v)

	t := triangularPrism{base: 10, attitude: 10, hieght: 10}
	VolumeOf(t)
	v = VolumeOf(t)
	fmt.Printf("%T %v\n", v, v)
}

// inside interface have type and value
type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	hieght   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

// จง implement "volumer" interfaceให้ type cube และ triangularPrism โดยที่
// volume ของ cube จะเท่ากับ edge * edge * edge
// และ volume ของ triangularPrism จะเท่ากับ 0.5 * base * attitude * height

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.hieght
}
