package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! 👋 from Go API")
}


func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Println("Server running on http://localhost:8080/hello")
    http.ListenAndServe(":8080", nil)
}
