package main

import (
    "net/http"
)

func main() {
    // Serve the HTML file at the root ("/")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    // Start the server on port 8080
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
