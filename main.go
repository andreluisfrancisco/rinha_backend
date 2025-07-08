package main

import (
    "fmt"
    "net/http"
    "time"
)

var client = &http.Client{
    Transport: &http.Transport{
        MaxIdleConns:        1000,
        MaxIdleConnsPerHost: 1000,
        IdleConnTimeout:     90 * time.Second,
    },
    Timeout: 2 * time.Second,
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Rinha backend")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server on :9999")
    http.ListenAndServe(":9999", nil)
}