package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// TravelsRoute Struct
type TravelsRoute struct{ *Route }

// index - Get all works
func (route *TravelsRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Travel{Entity: route.e}
	es := e.GetAll()
	route.serveGetAll(w, &es)
}

// byKey - Get one entity by its key
func (route *TravelsRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Travel{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}
