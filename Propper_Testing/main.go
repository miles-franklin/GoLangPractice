// Go: a simple programming environment
// Øredev Conference
// https://vimeo.com/53221558

package main

import (
	"fmt"
	"math"
)

type Vect struct {
	X, Y, Z float64
}

func (v Vect) GetMagnitude() float64{
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func main() {
	// var v = Vect{X: 10, Y: 11, Z: 12}
	// var v = Vect{X: 4, Y: 4, Z: 2}
	var v = Vect{X: 1, Y: 1, Z: 1}
	
	fmt.Println(v.GetMagnitude())
}