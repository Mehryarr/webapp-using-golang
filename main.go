package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldWriter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/myname":
		fmt.Fprint(w, "Mehryar")
	default:
		fmt.Fprint(w, "fatal error")
	}
	fmt.Printf("handling function with %s request \n", r.Method)

}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("html vs plain")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello, Welcome!</h1>")
}

func main() {
	http.HandleFunc("/", helloWorldWriter)
	http.HandleFunc("/htmlvsplain", htmlVsPlain)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
