package main

import "fmt"

func main() {
    b := 6 

    var bPtr *int // *int is used delcare variable
                   // b_ptr to be a pointer to an int

    bPtr = &b     // b_ptr is assigned value that is the address
                       // of where variable b is stored

    // Shorhand for the above two lines is:
    // b_ptr := &b

    fmt.Printf("address of b_ptr: %p\n", bPtr)

    // We can use *b_ptr get the value that is stored
    // at address b_ptr, or dereference the pointer 
    fmt.Printf("value stored at b_ptr: %d\n", *bPtr)

}

/*
Note that * can be used for two different things 
1) to declare a variable to be a pointer 
2) to dereference a pointer.

https://stackoverflow.com/questions/33242850/in-golang-what-is-the-difference-between-and
*/