package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var arr []int
	return func() int{
		if len(arr) == 0 {
			arr = append(arr, 0)
		} else if len(arr) == 1 {
			arr = append(arr, 1)
		} else {
			// T.I.L. that negative indexing doesn't exist over here in Go
			arr = append(arr, arr[len(arr)-1] + arr[len(arr)-2])
		}
		return arr[len(arr)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
