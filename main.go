package main

import (
	"context"
	"errors" // functions to manipulate errors
	"fmt"    // formatted IO
	"io"     // basic interface for IO
	"net"
	"net/http" // http client and server implementations
	"os"       // interface for operating system functionality
)

func main()  {

	mux := http.NewServeMux() // single connection between cilent and webserver
	// Multi threading

	mux.HandleFunc("/", getRoot)             // HANDLER FUNCTIONS
	mux.HandleFunc("/hello", getHello)       // HANDLER FUNCTIONS

	ctx, cancelCtx := context.WithCancel(context.Background())


	serverOne := &http.Server {
		Addr: ":3333",
		Handler: mux,
		BaseContext: func (l net.Listener) context.Context  {
			ctx := context.WithValue(ctx, keyServerAdd, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server {
		Addr: ":4444",
		Handler: mux,
		BaseContext: func (l net.Listener) context.Context  {
			ctx := context.WithValue(ctx, keyServerAdd, l.Addr().String())
			return ctx
		},
	}

	go func ()  {
		err:= serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server one is closed\n")
		} else if err != nil {
			fmt.Printf("Error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func ()  {
		err:= serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server two is closed\n")
		} else if err != nil {
			fmt.Printf("Error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()

	err := http.ListenAndServe(":3333", mux)  // START SERVER, LISTEN FOR REQUESTS

	if errors.Is(err, http.ErrServerClosed) {  // checking if server is getting closed
		fmt.Printf("Server Closed\n")
	} else if err != nil { // if there is an error shut the shit down
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1) // Closes server down
	}

}

const keyServerAdd = "serverAdd"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got /request\n", ctx.Value(keyServerAdd))
	io.WriteString(w, "This is my website\n") // Writes to page
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got: /hello request\n", ctx.Value(keyServerAdd))
	io.WriteString(w, "Hello, HTTP!") // Writes to page
}