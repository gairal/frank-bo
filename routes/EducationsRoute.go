package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// EducationsRoute Struct
type EducationsRoute struct{ *Route }

// index - Get all Educations
func (route *EducationsRoute) index(w http.ResponseWriter, r *http.Request) {
	edu := &models.Education{Entity: route.e}
	edus := edu.GetAll()

	route.serveGetAll(w, &edus)
}
