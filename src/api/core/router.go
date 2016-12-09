package core

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func BindRouter(router *mux.Router, base string, routes Routes) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		p := path.Join(base, route.Pattern)

		router.
			Methods(route.Method).
			Path(p).
			Name(route.Name).
			Handler(handler)

		if p != "/" {
			router.
				Methods(route.Method).
				Path(p + "/").
				Name(route.Name).
				Handler(handler)
		}
	}
}
