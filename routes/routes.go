package routes

import "net/http"

// Route - route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - all routes
type Routes []Route

var routes = Routes{
	Route{
		"WorksIndex",
		"GET",
		"/works",
		WorksIndex,
	},
}
