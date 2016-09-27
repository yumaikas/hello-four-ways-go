package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", root)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strings.Repeat("\n", 4))
	path := strings.Replace(r.URL.Path[1:], "+", " ", -1)
	fmt.Fprintf(w, "Hello, %v", path)
}
