package routes

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gairal/frank-gairal-bo/models"
)

// WorksRoute Struct
type WorksRoute struct {
	db *datastore.Client
}

// WorksIndex - Get all works
func (route *WorksRoute) worksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var work models.Work
	works := work.GetAll(route.db, "work")
	json.NewEncoder(w).Encode(works)
}
