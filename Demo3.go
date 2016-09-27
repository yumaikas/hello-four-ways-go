package main

import (
	"fmt"
	"mime"
	"net/http"
)

// SHOW OMIT
// Clickonce
func init() {
	mime.AddExtensionType(".application", "application/x-ms-application")
	mime.AddExtensionType(".manifest", "application/x-ms-manifest")
}

// https://www.junglecoder.com/downloads/areacodedialer/ as an example
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Listening on localhost:8080")
	// Error handling elided
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// ENDSHOW OMIT
