package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", htmlPlain)
	http.HandleFunc("/server", server)
	http.HandleFunc("/timeout", timeout)
	//http.ListenAndServe(":8080", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()

}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not timeout")

}

func server(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ninja":
		fmt.Fprint(w, "Allan")
	case "/web":
		fmt.Fprint(w, "Web request")
	default:
		fmt.Fprint(w, "Hello World")

	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Its me ")
	w.Header().Set("content type", "text/index.html")
	fmt.Fprintf(w, "<h1>Hello World <h2>")
}
