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

	// r.addRoute(RouteStructs{
	// 	RouteStruct{
	// 		"initDb",
	// 		"POST",
	// 		"/init",
	// 		r.init,
	// 	},
	// })
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

// init - Add one test entity
func (r *Router) init(w http.ResponseWriter, req *http.Request) {
	// 	ctx := appengine.NewContext(r)

	// 	k1 := datastore.NewKey(ctx, "image", "", 0, nil)
	// 	i1 := models.Image{
	// 		Name: "ekino",
	// 	}

	// 	ki1, err := datastore.Put(ctx, k1, &i1)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	k2 := datastore.NewKey(ctx, "image", "", 0, nil)
	// 	i2 := models.Image{
	// 		Name: "backelite",
	// 	}

	// 	ki2, err := datastore.Put(ctx, k2, &i2)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	e1 := models.Education{
	// 		ShortDescription: "Foo",
	// 		YearIn:           time.Now(),
	// 		YearOut:          time.Now(),
	// 		Diploma:          "osef",
	// 		Place:            "toto",
	// 		Name:             "toto",
	// 		Website:          "tata",
	// 		ImageKey:         ki1,
	// 	}
	// 	e2 := models.Education{
	// 		ShortDescription: "Bar",
	// 		YearIn:           time.Now(),
	// 		YearOut:          time.Now(),
	// 		Diploma:          "osef",
	// 		Place:            "tata",
	// 		Name:             "toto",
	// 		Website:          "tata",
	// 		ImageKey:         ki2,
	// 	}

	// 	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "education", nil), &e1); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "education", nil), &e2); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
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
