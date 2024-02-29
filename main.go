package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
    "log"
)

func handleAssetRequest(w http.ResponseWriter, r *http.Request) {

    buf, err := ioutil.ReadFile("assets/gobin_icon.png")

    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "image/png")
    w.Write(buf)
}

func main() {
	// Serve the HTML file at the root ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/assets/gobin-icon", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path)
		http.ServeFile(w, r, "assets/gobin_icon.png")
	})

	// Start the server on port 8080
	fmt.Println("Start the server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
