package main

import (
	"net/http"
)

func main() {
	c := cache.New("inmemory")
	http.ListenAndServe(":12345", nil)
}
