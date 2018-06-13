package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gairal/frank-gairal-bo/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router - Route object
type Router struct {
	routes RouteStructs
}

// addRoute - Add new route
func (r *Router) addRoute(routes RouteStructs) {
	r.routes = append(r.routes, routes...)
}

// getRoutes - Add new route
func (r *Router) getRoutes() {
	var entity models.IEntity

	entity = &models.Travel{}
	travels := &TravelsRoute{&Route{r, entity}}
	r.addRoute(travels.getRoutes())

	entity = &models.Image{}
	images := &ImagesRoute{&Route{r, entity}}
	r.addRoute(images.getRoutes())

	entity = &models.Education{}
	edus := &EducationsRoute{&Route{r, entity}}
	r.addRoute(edus.getRoutes())
}

// RegisterHandlers - instanciate the router
func RegisterHandlers() {
	r := &Router{}
	r.getRoutes()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range r.routes {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))
}

// serveGetAll - Send JSON content
func (r *Router) serve(w http.ResponseWriter, entities interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(entities)
}

// getKey - getKey param
func (r *Router) getKey(req *http.Request) int64 {
	keyStr := mux.Vars(req)["key"]
	key, _ := strconv.ParseInt(keyStr, 10, 64)

	return key
}

// return RouteStructs{
// RouteStruct{
// 	"worksIndex",
// 	"GET",
// 	"/works",
// 	works.index,
// },
// RouteStruct{
// 	"worksByKey",
// 	"GET",
// 	"/works/{key}",
// 	works.byKey,
// },
// RouteStruct{
// 	"educationsIndex",
// 	"GET",
// 	"/educations",
// 	educations.index,
// },
// RouteStruct{
// 	"educationsByKey",
// 	"GET",
// 	"/educations/{key}",
// 	educations.byKey,
// },
// RouteStruct{
// 	"skillsIndex",
// 	"GET",
// 	"/skills",
// 	skills.index,
// },
// RouteStruct{
// 	"skillsByCategory",
// 	"GET",
// 	"/skills/categories",
// 	skills.byCategory,
// },
// RouteStruct{
// 	"skillsByKey",
// 	"GET",
// 	"/skills/{key}",
// 	skills.byKey,
// },
// RouteStruct{
// 	"interestsIndex",
// 	"GET",
// 	"/interests",
// 	interests.index,
// },
// RouteStruct{
// 	"interestsByCategory",
// 	"GET",
// 	"/interests/categories",
// 	interests.byCategory,
// },
// RouteStruct{
// 	"interestsByKey",
// 	"GET",
// 	"/interests/{key}",
// 	interests.byKey,
// },
// RouteStruct{
// 	"travelsIndex",
// 	"GET",
// 	"/travels",
// 	travels.index,
// },
// RouteStruct{
// 	"travelsByKey",
// 	"GET",
// 	"/travels/{key}",
// 	travels.byKey,
// },
// RouteStruct{
// 	"travelsAdd",
// 	"POST",
// 	"/travels",
// 	travels.add,
// },
// }
