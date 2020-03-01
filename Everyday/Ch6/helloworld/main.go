package main

import "net/http"

import "fmt"

import "log"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, 세계!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}