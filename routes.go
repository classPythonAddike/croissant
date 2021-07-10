package croissant

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Route struct {
	Path   string
	Get    func(http.ResponseWriter, *http.Request) Response
	Post   func(http.ResponseWriter, *http.Request) Response
	Delete func(http.ResponseWriter, *http.Request) Response
	Patch  func(http.ResponseWriter, *http.Request) Response
}

func WriteResponse(
	handler func(http.ResponseWriter, *http.Request) Response,
	w http.ResponseWriter,
	resp *http.Request,
) {

	var response Response = handler(w, resp)

	w.WriteHeader(response.StatusCode)

	if response.Html != "" {
		fmt.Fprint(w, response.Html)
	} else if response.Json != nil {
		json.NewEncoder(w).Encode(response.Json)
	}
}

func (r *Route) RequestHandler(w http.ResponseWriter, resp *http.Request) {
	var method string = resp.Method

	switch {

	case (method == "GET" || method == ""):
		if r.Get != nil {
			WriteResponse(r.Get, w, resp)
		} else {
			w.WriteHeader(405)
		}

	case method == "POST":
		if r.Post != nil {
			WriteResponse(r.Post, w, resp)
		} else {
			w.WriteHeader(405)
		}
	case method == "DELETE":
		if r.Delete != nil {
			WriteResponse(r.Delete, w, resp)
		} else {
			w.WriteHeader(405)
		}
	case method == "PATCH":
		if r.Patch != nil {
			WriteResponse(r.Patch, w, resp)
		} else {
			w.WriteHeader(405)
		}
	}
}
