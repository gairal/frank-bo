package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// WorksRoute Struct
type WorksRoute struct{ *Route }

// index - Get all works
func (route *WorksRoute) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	e := &models.Entity{DbClient: route.db}
	works := e.GetAll("work")
	json.NewEncoder(w).Encode(works)
}
