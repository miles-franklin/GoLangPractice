package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method
// The leading reciever argument `(v Vertex)` is what makes this a method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) AddOne(){
	v.X++
	v.Y++
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))		// Function Call
	fmt.Println(v.Abs())	// Method Call

	
	fmt.Printf("Before `v.AddOne()`:\t%v\n", v)
	v.AddOne()
	fmt.Printf("After `v.AddOne()`:\t%v\n", v)
}