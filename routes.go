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

	rawReq.ParseForm()

	for param, paramType := range r.Expects {
		valid := ValidTypes[paramType](rawReq.FormValue(param))
		if !valid {
			w.WriteHeader(400)
			log.Printf("%v - %v returned %v", method, r.Path, 400)
			return
		}
	}

	methodMap := map[string](func(*Request, FormBody) Response){
		"GET":    r.Get,
		"":       r.Get,
		"POST":   r.Post,
		"PATCH":  r.Patch,
		"DELETE": r.Delete,
	}

	if methodMap[method] != nil {
		writeResponse(method, r, methodMap[method], FormBody(rawReq.Form), w, req)
	} else {
		w.WriteHeader(405)
		log.Printf("%v - %v returned %v", method, r.Path, 405)
	}
}
