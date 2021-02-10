package main

import (
	"net/http"

	"example.com/just-a-test/src/env"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":"+env.Port(), nil)
}
