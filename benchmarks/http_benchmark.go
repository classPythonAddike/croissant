package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func TestHttp() {
	http.HandleFunc("/", GetResponse)
	http.ListenAndServe(":8080", nil)
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	num, err := strconv.Atoi(r.Form.Get("integer"))

	if err != nil {
		fmt.Fprint(w, "Invalid type for form parameter integer - expected int!")
		return
	}

	if num%2 != 0 {
		w.WriteHeader(200)
		fmt.Fprint(w, fmt.Sprintf("Error - Odd number given - %v", num))
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, fmt.Sprintf("<h1>Even Number given - %v</h1>", num))
	}
}
