package routes

import (
	"net/http"
)

// EducationsRoute Struct
type EducationsRoute struct{ *Route }

// index - Get all works
func (route *EducationsRoute) index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// e := &models.Entity{DbClient: route.db}
	// // edus := e.GetAll(models.Education, "education")
	// // json.NewEncoder(w).Encode(edus)
}
