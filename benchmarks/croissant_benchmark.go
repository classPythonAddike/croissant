package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/classPythonAddike/croissant"
)

func TestCroissant() {
	router := croissant.Router{}

	router.InitRoutes(1)                                          // Only 1 route
	router.AddRoute(croissant.Route{Path: "/", Get: GetHomePage}) // Only allows Get Requests

	router.Serve()
}

func GetHomePage(w http.ResponseWriter, r *http.Request) croissant.Response {
	num := rand.Intn(100)

	if num%2 != 0 {
		return croissant.Response{
			Html:       fmt.Sprintf("Odd number, error ocurred - %v", num),
			StatusCode: 400,
		}
	}

	return croissant.Response{
		Html:       fmt.Sprintf("<h1>Even Number - %v</h1>", num),
		StatusCode: 200,
	}
}
