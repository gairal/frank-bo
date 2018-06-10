package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
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

// Route - Route object
type Route struct {
	e *models.Entity
}

// GetRoutes - List all routes
func (r *Route) getRoutes() RouteStructs {
	works := &WorksRoute{r}
	educations := &EducationsRoute{r}

	return RouteStructs{
		RouteStruct{
			"worksIndex",
			"GET",
			"/works",
			works.index,
		},
		RouteStruct{
			"educationsIndex",
			"GET",
			"/educations",
			educations.index,
		},
	}
}

// serveGetAll - Send JSON content
func (r *Route) serveGetAll(w http.ResponseWriter, entities *interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(entities)
}
