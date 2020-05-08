package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
add a really simple authentication decorator function that will check to see if the Authorized header is set to true on the incoming request.
*/

/*
takes in a function that matches the same signature as our original homePage function. This then returns a http.Handler.
*/
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking Authorized header set...")

		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!")
			fmt.Fprintf(w, "Not Authorized!")
		}
	})
}

//up a net/http router that serves a single / endpoint.
func homePage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Authorized", "true")
	r.Header.Set("Authorized", "true")
	fmt.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to the Homepage")
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

/*
The key thing to note however, is the fact that we’ve managed to decorate an existing endpoint and add some form of authentication around said endpoint without having to alter the existing implementation of that function.

This highlights the key benefits of the decorator pattern, where wrapping code within our codebase is incredibly simple. We can easily add new authenticated endpoints using this same method
*/
