package main

import (
	"fmt"
	"math"
)

// IVertex interface operations
type IVertex interface {
	Abs() float64
	Scale(f float64)
}

// Vertex defines a corner where lines meet
type Vertex struct {
	X, Y float64
}

// Abs return the absolute value of a Vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale modifies a vertex for enlarging or reducing it
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var v IVertex = &Vertex{3, 4}
	describe(v)
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v.Abs())

	var v2 interface{} = &Vertex{5, 6}
	if vertex, ok := v2.(IVertex); ok {
		describe(vertex)
		fmt.Println(vertex.Abs())
		vertex.Scale(2)
		fmt.Println(vertex.Abs())
	}
}
