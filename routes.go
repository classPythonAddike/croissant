package croissant

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Request http.Request
type FormBody url.Values

/*
Route class

Each route has attributes for get, post, patch and delete requests.
If an attribute is nil, it is assumed that the method is not allowed for the route.

Path - string which identifies this route. Example: "/route/id" means that
this route's methods should be invoked for the "/route/id" path

Expects - map of string-string, which defines the form body parameters expected, and
their types. Valid types - "int", "float", "bool", "string"
*/
type Route struct {
	Path    string
	Expects map[string]string

	Get    func(*Request, FormBody) Response
	Post   func(*Request, FormBody) Response
	Delete func(*Request, FormBody) Response
	Patch  func(*Request, FormBody) Response
}

func writeResponse(
	method string,
	r *Route,
	handler func(*Request, FormBody) Response,
	body FormBody,
	w http.ResponseWriter,
	resp Request,
) {

	var response Response = handler(&resp, body)

	if response.StatusCode == 0 { // status code not provided
		response.StatusCode = 200
	}

	w.WriteHeader(response.StatusCode)
	log.Printf("%v - %v returned %v", method, r.Path, response.StatusCode)

	if response.Html != "" {
		fmt.Fprint(w, response.Html)
	} else if response.Json != nil {
		json.NewEncoder(w).Encode(response.Json)
	}
}

func (r *Route) requestHandler(w http.ResponseWriter, rawReq *http.Request) {

	req := Request(*rawReq)
	var method string = req.Method

	methodMap := map[string](func(*Request, FormBody) Response){
		"GET":    r.Get,
		"":       r.Get,
		"POST":   r.Post,
		"PATCH":  r.Patch,
		"DELETE": r.Delete,
	}

	// Check if the route has a handler for the method
	if methodMap[method] != nil {

		rawReq.ParseForm()

		for param, paramType := range r.Expects {

			valid := validTypes[paramType](rawReq.FormValue(param)) // Check parameter type

			if !valid {
				w.WriteHeader(400) // Invalid type for a form parameter

				msg := fmt.Sprintf(
					"Invalid type for form parameter %v - expected %v",
					param,
					paramType,
				)

				fmt.Fprint(w, msg)
				log.Printf("%v - %v returned %v (%v)", method, r.Path, 400, msg)
				return
			}
		}

		writeResponse(method, r, methodMap[method], FormBody(rawReq.Form), w, req)

	} else {
		w.WriteHeader(405) // Method not provided
		fmt.Fprint(
			w,
			"Method Not Allowed!",
		)
		log.Printf("%v - %v returned %v (Method Not Allowed)", method, r.Path, 405)
	}
}
