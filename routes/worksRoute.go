package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// WorksRoute Struct
type WorksRoute struct{ *Route }

// index - Get all works
func (route *WorksRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Work{Entity: route.e}
	es := e.GetAll()
	route.serveGetAll(w, &es)
}

// byKey - Get one entity by its key
func (route *WorksRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Work{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}
