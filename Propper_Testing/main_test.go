// go test
// go test -v
// go test -test.bench=GetMagnitude -v

package main

import (
	"testing"
	"fmt"
)

func TestGetMagnitude(t *testing.T) {
	var tests = []struct {
		input Vect
		expected float64
	} {
		{Vect{4,4,2}, float64(6)},
		{Vect{1,1,1}, float64(1.7320508075688772)},
		{Vect{10,11,12}, float64(19.1049731745428)},
	}

	for _, test := range tests {
		if test.input.GetMagnitude() != test.expected {
			t.Errorf("Vect(%v,%v,%v) = %v. Was expecting %v.",
					test.input.X, test.input.Y, test.input.Z, test.input.GetMagnitude(), test.expected)
		}
	}

}

func BechmarkGetMagnitude(b *testing.B) {
	v := Vect{1,1,1}
	for i := 0; i < b.N; i++ {
		v.GetMagnitude()
	}
}

func ExampleGetMagnitude() {
	v1 := Vect{4,4,2}
	fmt.Println(v1.GetMagnitude())

	// Output:
	// 6
}