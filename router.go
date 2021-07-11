package croissant

import (
	"fmt"
	"net/http"
)

type Router struct {
	Routes map[string]Route
}

func NewRouter(NumRoutes int) Router {
	r := Router{}
	r.Routes = make(map[string]Route, NumRoutes)
	return r
}

func (r *Router) AddRoute(route Route) {

	var isValid bool

	for _, paramType := range route.Expects {
		isValid = false
		for validType, _ := range ValidTypes {
			if paramType == validType {
				isValid = true
				break
			}
		}

		if !isValid {
			panic(fmt.Sprintf("Unsupported type for form body expectation - %v", paramType))
		}
	}

	http.HandleFunc(route.Path, route.requestHandler)
	r.Routes[route.Path] = route
}

func (r *Router) Serve(host string) {
	http.ListenAndServe(host, nil)
}
