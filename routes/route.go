package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
	"google.golang.org/appengine"
)

// RouteStruct - route structure
type RouteStruct struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// RouteStructs - all routes
type RouteStructs []RouteStruct

// IRoute - Route interface
type IRoute interface {
	getRoutes() RouteStructs
	index(w http.ResponseWriter, r *http.Request)
	byKey(w http.ResponseWriter, r *http.Request)
}

// Route - Route object
type Route struct {
	router *Router
	entity models.IEntity
}

// index - Get all entities
func (route *Route) index(w http.ResponseWriter, r *http.Request) {
	es := route.entity.GetAll(appengine.NewContext(r))
	route.router.serve(w, &es)
}

// byKey - Get one entity by its key
func (route *Route) byKey(w http.ResponseWriter, r *http.Request) {
	route.entity.Get(appengine.NewContext(r), route.router.getKey(r))
	route.router.serve(w, route.entity)
}
