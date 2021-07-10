package croissant

import (
	"fmt"
	"net/http"
)

func HelloWorld() {
	fmt.Println("Hello, World")
}

type Router struct {
	Routes map[string]Route
}

func (r *Router) InitRoutes(NumRoutes int) {
	r.Routes = make(map[string]Route, NumRoutes)
}

func (r *Router) AddRoute(route Route) {
	http.HandleFunc(route.Path, route.RequestHandler)
	r.Routes[route.Path] = route
}

func (r *Router) Serve() {
	http.ListenAndServe(":8080", nil)
}
