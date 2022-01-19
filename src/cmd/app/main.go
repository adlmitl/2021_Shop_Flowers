package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Flowers")
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)

	http.ListenAndServe(":8000", mux)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Flowers"))
}
