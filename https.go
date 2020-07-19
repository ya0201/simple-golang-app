package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServeTLS(":10443", "self-signed.crt", "self-signed.key", nil)
}
