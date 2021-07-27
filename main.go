package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func addPage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hi")
}

func main() {
	fs := http.FileServer(http.Dir("./wasm"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", addPage)
	log.Print("Serving " + "./wasm" + " on http://localhost:8080")
	http.ListenAndServe(":8080", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		fs.ServeHTTP(resp, req)
	}))
}
