package main

import (
	"fmt"
	"mime"
	"net/http"
)

var q = mime.WordDecoder{}

// SHOW OMIT
func init() {
	mime.AddExtensionType(".image", "image/svg+xml")
}

// https://www.junglecoder.com/downloads/areacodedialer/ as an example
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Listening on localhost:8080")
	// Error handling elided
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// ENDSHOW OMIT
