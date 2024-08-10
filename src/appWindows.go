// Copyright 2024 The Author. All rights reserved.

// Web server in 3 lines.

package main

import (
	"log"
	"net/http"
)

// Handler serves Outputs text.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}
func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
