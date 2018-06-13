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
func (r *Route) index(w http.ResponseWriter, req *http.Request) {
	es := r.entity.GetAll(appengine.NewContext(req))
	r.router.serve(w, &es)
}

// byKey - Get one entity by its key
func (r *Route) byKey(w http.ResponseWriter, req *http.Request) {
	r.entity.Get(appengine.NewContext(req), r.router.getKey(req))
	r.router.serve(w, r.entity)
}
