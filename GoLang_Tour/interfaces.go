package main

import (
	"fmt"
)

// Still not sure how to use this...
type Person interface{
	Introduce()
}

type Student struct{
	// Grades map[string]string
	Name string
	GPA float64
}

type Teacher struct{
	Name string
	Students []*Student
}

func (t *Teacher)Introduce() {
	fmt.Printf("Hello, my name is %v\n", (*t).Name)
}

func (s *Student)Introduce() {
	fmt.Printf("Hello, my name is %v\n", (*s).Name)
}

// Stringers - this is how we define the way we want fmt.Println() to handle printing out object
func (s Student) String() string {
	return fmt.Sprintf("%v (%.2f GPA)", s.Name, s.GPA)
}

func main(){
	EmptyInterfaces()
	TypeSwitches()

	var t Teacher
	t.Name = "Proj G"

	var s1 Student
	s1.Name = "Miles"
	s1.GPA = 4.0
	t.Students = append(t.Students, &s1)

	var s2 Student
	s2.Name = "Julian"
	s2.GPA = 3.0
	t.Students = append(t.Students, &s2)

	t.Introduce()
	fmt.Println()

	for _, s := range t.Students {
		s.Introduce()
		fmt.Println(*s)
		fmt.Println()

	}
}

func EmptyInterfaces(){
	fmt.Println("EmptyInterfaces()")
	defer fmt.Println("")

	// This technique can be used to create a variable that can hold any type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Assertions(){
	fmt.Println("Assertions()")
	defer fmt.Println("")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// The following line will cause a runtime panic unless the the boolean 
	// ok is captured as well
	// f = i.(float64) // panic
	// fmt.Println(f)
}

func do(i interface{}) {
	// By taking in a generic type (an `interface {}`), we can swtich based on
	// which type is passed through
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeSwitches() {
	fmt.Println("TypeSwitches()")
	defer fmt.Println("")

	do(21)
	do("hello")
	do(true)
}

