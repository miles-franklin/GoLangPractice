//	Description:
//		`boring()` is an infinite loop that adds to adds the the channel after
//		some variable number of seconds. If that delay it too long
//		('<-time.After(1 * time.Second)'), the program will exit
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
			sleepTime := time.Duration(2*rand.Intn(750)) * time.Millisecond
			time.Sleep(sleepTime)
		}
	}()
	return c
}

func testBoring(){
	fmt.Println("testBoring():")
	defer fmt.Println("Done\n")

	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func boringQuit(name string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := rand.Intn(10); i>=0; i-- {
			c <- fmt.Sprintf("%s %d", name, i)
			sleepTime := time.Duration(2*rand.Intn(750)) * time.Millisecond
			time.Sleep(sleepTime)
		}
		quit <- "Time to exit"
		fmt.Printf("%s %v\n", name, <-quit) // TODO: Not sure why its not getting here...
	}()
	return c
}

func testBoringQuit(){
	fmt.Println("testBoringQuit():")
	defer fmt.Println("Done\n")

	quit := make(chan string)
	c := boringQuit("Joe", quit)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			// Cleanup
			quit <- "main() got the exit message"
			return
		}
	}
}

func main() {
	testBoring()
	testBoringQuit()
	
}