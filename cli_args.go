package main

import (
	"fmt"
	"flag"
)

var (
	msg = flag.String("msg", "hello world", "Give a test message to print out.")
)

func main() {
	flag.Parse()
	fmt.Println(*msg)
}