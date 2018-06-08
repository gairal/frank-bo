package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// RegisterHandlers - instanciate the router
func RegisterHandlers() {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		// handler = tools.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}
