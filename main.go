package main

import (
	"fmt"
	"net/http"
)

func helloWorldWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorldWriter)
	http.ListenAndServe("", nil)
}
