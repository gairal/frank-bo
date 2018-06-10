package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// WorksRoute Struct
type WorksRoute struct{ *Route }

// index - Get all works
func (route *WorksRoute) index(w http.ResponseWriter, r *http.Request) {
	work := &models.Work{Entity: route.e}
	works := work.GetAll()
	route.serveGetAll(w, &works)
}
