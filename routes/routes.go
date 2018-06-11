package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gairal/frank-gairal-bo/models"
	"github.com/gorilla/mux"
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
	skills := &SkillsRoute{r}
	interests := &InterestsRoute{r}
	travels := &TravelsRoute{r}

	return RouteStructs{
		RouteStruct{
			"worksIndex",
			"GET",
			"/works",
			works.index,
		},
		RouteStruct{
			"worksByKey",
			"GET",
			"/works/{key}",
			works.byKey,
		},
		RouteStruct{
			"educationsIndex",
			"GET",
			"/educations",
			educations.index,
		},
		RouteStruct{
			"educationsByKey",
			"GET",
			"/educations/{key}",
			educations.byKey,
		},
		RouteStruct{
			"skillsIndex",
			"GET",
			"/skills",
			skills.index,
		},
		RouteStruct{
			"skillsByCategory",
			"GET",
			"/skills/categories",
			skills.byCategory,
		},
		RouteStruct{
			"skillsByKey",
			"GET",
			"/skills/{key}",
			skills.byKey,
		},
		RouteStruct{
			"interestsIndex",
			"GET",
			"/interests",
			interests.index,
		},
		RouteStruct{
			"interestsByCategory",
			"GET",
			"/interests/categories",
			interests.byCategory,
		},
		RouteStruct{
			"interestsByKey",
			"GET",
			"/interests/{key}",
			interests.byKey,
		},
		RouteStruct{
			"travelsIndex",
			"GET",
			"/travels",
			travels.index,
		},
		RouteStruct{
			"travelsByKey",
			"GET",
			"/travels/{key}",
			travels.byKey,
		},
	}
}

// serveGetAll - Send JSON content
func (r *Route) serveGetAll(w http.ResponseWriter, entities interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(entities)
}

// getKey - getKey param
func (r *Route) getKey(req *http.Request) int64 {
	keyStr := mux.Vars(req)["key"]
	key, _ := strconv.ParseInt(keyStr, 10, 64)

	return key
}
