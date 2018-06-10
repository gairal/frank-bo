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
	for _, route := range GetRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		// handler = tools.Logger(handler, route.Name)
		// handler = BindDb(route.HandlerFunc, db)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}

// BindDb - bind DB client to the handler
// func BindDb(inner http.Handler, db *datastore.Client) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		inner.ServeHTTP(w, r, db)
// 	})
// }
