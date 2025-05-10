package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routes_exist(t *testing.T) {
	testApp := Config {}
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router) // casting to chi.Router
	routes := []string{"/authenticate"}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, expectedRoute string) {
	found := false
	_ = chi.Walk(routes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if expectedRoute == route {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not find %s in registered routes", expectedRoute)
	}
}