package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from Go and k8s")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
