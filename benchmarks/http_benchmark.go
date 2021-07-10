package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func TestHttp() {
	http.HandleFunc("/", GetResponse)
	http.ListenAndServe(":8080", nil)
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(100)

	if num%2 != 0 {
		w.WriteHeader(500)
		fmt.Fprint(w, fmt.Sprintf("Error - Odd number given - %v", num))
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, fmt.Sprintf("<h1>Even Number given - %v</h1>", num))
	}
}
