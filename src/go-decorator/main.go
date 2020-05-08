package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

// takes a function as a parameter
func coolFunc(a func()) {
	fmt.Printf("Start time: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

/*
we’ve been able to effectively wrap my original function without having to alter it’s implementation
*/

func main() {
	//print out the type of the value we pass in as our second argument.
	fmt.Printf("Type: %T\n", myFunc)
	coolFunc(myFunc)
}

/**
In Go, functions are deemed as first class objects which essentially means you can pass them around just as you would a variable.

**/