package main

import (
	"fmt"
	"net/http"
)

const port = ":8888"

func main() {

	http.HandleFunc("/", Hi)
	fmt.Println("listening", port)
	http.ListenAndServe(port, nil)
}

func Hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hihi")
}
