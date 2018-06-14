package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gairal/frank-gairal-bo/mocks"
	"github.com/gairal/frank-gairal-bo/models"
	"github.com/rs/cors"
	"google.golang.org/appengine"

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
	var route IRoute

	entity = &models.Travel{}
	route = &TravelsRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Image{}
	route = &ImagesRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Category{}
	route = &CategoriesRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Education{}
	route = &EducationsRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Work{}
	route = &WorksRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Skill{}
	route = &SkillsRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	entity = &models.Interest{}
	route = &InterestsRoute{&Route{r, entity}}
	r.addRoute(route.getRoutes())

	// TODO: remove
	r.addRoute(RouteStructs{
		RouteStruct{
			"initDb",
			"POST",
			"/init",
			r.init,
		},
	})
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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://frank.gairal.com", "http://frank.gairal.com", "http://localhost:3000"},
	})

	handler := c.Handler(router)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, handler))
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

// init - Add one test entity
func (r *Router) init(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	m := &mocks.Mock{}

	m.InitDB(ctx)
}
