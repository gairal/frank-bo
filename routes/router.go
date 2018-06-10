package routes

import (
	"net/http"
	"os"

	"github.com/gairal/frank-gairal-bo/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// RegisterHandlers - instanciate the router
func RegisterHandlers() {
	r := &Route{e: models.InitDbClient()}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range r.getRoutes() {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}
