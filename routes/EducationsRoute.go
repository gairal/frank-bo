package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// EducationsRoute Struct
type EducationsRoute struct{ *Route }

// index - Get all Educations
func (route *EducationsRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Education{Entity: route.e}
	es := e.GetAll()
	route.serveGetAll(w, &es)
}

// byKey - Get one entity by its key
func (route *EducationsRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Education{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}
