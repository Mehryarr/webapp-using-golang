package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	//w.Header().Set("connection", "keep-alive")
	//w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Hello, Welcome!</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "did not timeout")
}

func main() {
	http.HandleFunc("/", helloWorldWriter)
	http.HandleFunc("/htmlvsplain", htmlVsPlain)
	http.HandleFunc("/timeout", timeout)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	server := http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	log.Fatal(server.ListenAndServe())
}
