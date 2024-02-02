package main

import "fmt"

func adder() func(int) int {
	// From the tutorial
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func print() func () string {
	// This is my first attempt at writing a closure funciton. The expected behaviour
	// is that calling this function each time will pop the first variable from
	// the name slice and append it to the out put string. So:
	// f := print()
	// 
	// f() -> Miles
	// f() -> Miles Julian
	// f() -> Miles Julian Franklin
	// f() -> Miles Julian Franklin

	name := []string{"Miles", "Julian", "Franklin"}
	var str string 
	return func() string {
		if len(name) > 0{
			x := pop(&name)
			str = str + " " + x
		}
		return str
	}
}

func pop[T any](s *[]T) T {
	// My helper function (bonus practice for using pointers)
	// This function takes in the address of a slice of anny type and pops off
	// the first value, while manipulating the original slice being referenced.
	if len(*s) > 0 {
		first := (*s)[0]
		*s = (*s)[1:]
		return first
	} else {
		var empty T
		return empty
	}
}

func testPop(){
	// Testing pop() funciton
	arr := []int{1,2,3} 
	fmt.Println(arr)
	e := pop(&arr)
	fmt.Println(e, arr)
}

func testMyClosureFunc(){
	// Testing print() funciton
    var fSlice []func() string
    fSlice = append(fSlice, print())
    fSlice = append(fSlice, print())
	for _, f := range fSlice{
		for i := 0; i<4; i++{
			fmt.Println(f())
		}
		fmt.Println()
	}
}

func main() {
	// This was the original code given from https://go.dev/tour/moretypes/25
	pos, neg := adder(), adder()
	for i := 0; i < 0; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	
	testPop() 				// Works!
	testMyClosureFunc() 	// Works!
}
