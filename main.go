package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, 世界,World,Мир")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running simple-webapp app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
