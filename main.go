package main

import (
	"errors"   // functions to manipulate errors
	"fmt"      // formatted IO
	"io"       // basic interface for IO
	"net/http" // http client and server implementations
	"os"       // interface for operating system functionality
)

func main()  {
	http.HandleFunc("/", getRoot)             // HANDLER FUNCTIONS
	http.HandleFunc("/hello", getHello)       // HANDLER FUNCTIONS

	err := http.ListenAndServe(":3333", nil)  // START SERVER, LISTEN FOR REQUESTS

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil { // if there is an error shut the shit down
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1) // Closes server down
	}

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / Request")
	io.WriteString(w, "This is my website\n") // Writes to page
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /hello request")
	io.WriteString(w, "Hello, HTTP!") // Writes to page
}