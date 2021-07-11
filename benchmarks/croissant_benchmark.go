package main

import (
	"fmt"
	"strconv"

	"github.com/classPythonAddike/croissant"
)

func TestCroissant() {
	router := croissant.NewRouter(1) // Only 1 route
	router.AddRoute(
		croissant.Route{
			Path:    "/",
			Get:     GetHomePage,
			Expects: map[string]string{"integer": "int"},
		},
	) // Only allows Get Requests

	router.Serve(":8080")
}

func GetHomePage(r *croissant.Request, body croissant.FormBody) croissant.Response {
	num, _ := strconv.Atoi(body["integer"][0])

	if num%2 != 0 {
		return croissant.Response{
			Html:       fmt.Sprintf("Odd number, error ocurred - %v", num),
			StatusCode: 405,
		}
	}

	return croissant.Response{
		Html: fmt.Sprintf("<h1>Even Number - %v</h1>", num),
	}
}
