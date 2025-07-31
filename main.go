package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Hello from Go in Kubernetes! Host: %s\n", host)
}

func main() {
    http.HandleFunc("/", handler)
    port := "8080"
    log.Printf("Starting server on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
