package main

import (
	"fmt"
	"naosuke-ito-golang-sample-app/hello"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, hello.SayHello())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
