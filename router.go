package croissant

import (
	"fmt"
	"log"
	"net/http"
)

/*
croissant's Router class

A Router is a very simple struct - it only holds a map of the
routes that the api supports. To create a new router,
use `myRouter := croissant.NewRouter(num)`, where
`num` is the number of routes which will be added.
To add a new route, use `myRouter.AddRoute(route)`.
*/
type Router struct {
	Routes      map[string]Route
	routeLength int
}

/*
Function to return a new Router.
Provide the number of routes that will be defined as a parameter
*/
func NewRouter(NumRoutes int) Router {
	r := Router{routeLength: NumRoutes}
	r.Routes = make(map[string]Route, NumRoutes)
	return r
}

/*
Add a route to a router, providing the route struct
*/
func (r *Router) AddRoute(route Route) {

	var isValid bool

	if route.Expects != nil {
		for _, paramType := range route.Expects {
			isValid = false
			for validType := range validTypes {
				if paramType == validType {
					isValid = true
					break
				}
			}

			if !isValid {
				panic(fmt.Sprintf("Unsupported type for form body expectation - %v", paramType))
			}
		}
	}

	http.HandleFunc(route.Path, route.requestHandler)
	r.Routes[route.Path] = route
}

/*
Run the api, and listen on the provided port
*/
func (r *Router) Serve(host string) {

	if len(r.Routes) != r.routeLength {
		log.Fatalf(
			"Router was initialised with %v routes, but got %v routes",
			r.routeLength,
			len(r.Routes),
		)
	}

	http.ListenAndServe(host, nil)
}
