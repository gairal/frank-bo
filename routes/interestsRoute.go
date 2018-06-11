package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// InterestsRoute Struct
type InterestsRoute struct{ *Route }

// index - Get all skills
func (route *InterestsRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Interest{Entity: route.e}
	es := e.GetAll()
	route.serveGetAll(w, &es)
}

// byCategory - Get all skills by categories
func (route *InterestsRoute) byCategory(w http.ResponseWriter, r *http.Request) {
	e := &models.Interest{Entity: route.e}
	es := e.GetAllByCategory()
	route.serveGetAll(w, &es)
}

// byKey - Get one entity by its key
func (route *InterestsRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Interest{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}
