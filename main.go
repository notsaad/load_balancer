package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Fprint is for file/writer output
    fmt.Fprintf(w, "Load balancer is running!")
}

func main() {
    port := 8080
    http.HandleFunc("/", handler)

    log.Printf("Load Balancer running on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
