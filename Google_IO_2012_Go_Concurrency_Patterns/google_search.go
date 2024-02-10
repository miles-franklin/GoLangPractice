package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"reflect"
    "encoding/json"
	"io/ioutil"
	"os"
)

var (
	Web = fakeSearch("web")
	Img = fakeSearch("img")
	Vid = fakeSearch("vid")
)

type Search func(query string) Result

type Result string

func Google_0(query string) Result{
	// Proof of concept for basic Google()
	f := fakeSearch(query) // returns a Search()
	return f(query)
}

func Google_1(query string) (results []Result){
	// Grabs each search result in series
	results = append(results, Web(query))
	results = append(results, Img(query))
	results = append(results, Vid(query))
	return
}

func Google_2(query string) (results []Result){
	// Grabs each search result concurrently
	// Now, we are only waiting on the slowest search
	// Note that this is a "Fan In" Go pattern
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Img(query) }()
	go func() { c <- Vid(query) }()

	for i:=0; i<3; i++{
		results = append(results, <-c)
	}

	return
}

func Google_3(query string) (results []Result){
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Img(query) }()
	go func() { c <- Vid(query) }()

	// After 80ms, the program will add to the timeout channel. The select case
	// will see that there is data ready on the timeout channel and the execute
	// that case accordingly.
	timeout := time.After(80 * time.Millisecond)
	for i:=0; i<3; i++{
		select {
		case r := <-c:
			results = append(results, r)
		case <- timeout:
			fmt.Println("Timed out")
			return
		}
	}

	return
}

func Google(query string) (results []Result){
	c := make(chan Result)
	go func() { c <- First(query, Web, Web) }()
	go func() { c <- First(query, Img, Img) }()
	go func() { c <- First(query, Vid, Vid) }()

	timeout := time.After(80 * time.Millisecond)
	for i:=0; i<3; i++{
		select {
		case r := <-c:
			results = append(results, r)
		case <- timeout:
			fmt.Println("Timed out")
			return
		}
	}

	return
}

func TestGoogle(){
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func First(query string, replicas ...Search) Result  {
	// This funciton creates replica functions and return after the fastest
	// one pushes a response onto the channel
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func FirstSlice(query string, replicas []Search) (Result, time.Duration)  {
	// This funciton creates replica functions and return after the fastest
	// one pushes a response onto the channel
	start := time.Now()
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c, time.Since(start)
}

// TODO: Restructure this to record multiple tests at each count. In the
// analyiss 
func TestFirst(i int){
	results := make(map[int]string)
	for count:=1; count<=i; count++{
		fmt.Println(strings.Repeat("=", 50))
		fmt.Printf("Count = %v\n", count)
		arr := make([] Search, i)
		for i, _ := range arr{
			name := fmt.Sprintf("R%v", i)
			arr[i] = fakeSearch(name)
		}

		_, dur := FirstSlice("golang", arr)
		fmt.Println(reflect.TypeOf(dur))
		fmt.Printf("%v\n\n", dur)
		results[count] = fmt.Sprintf("%s", dur)
	}

	// Save to JSON file
	jsonString, err:= json.MarshalIndent(results, "", "    ")
	if err == nil {
		ioutil.WriteFile("Google_IO_2012_Go_Concurrency_Patterns/output.json", jsonString, os.ModePerm)
	} else {
		fmt.Println(err)
	}
}


func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	// TestGoogle()
	TestFirst(10)
}