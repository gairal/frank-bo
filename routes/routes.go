package routes

import (
	"net/http"

	"cloud.google.com/go/datastore"
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
	DbClient *datastore.Client
}

// GetRoutes - List all routes
func GetRoutes() RouteStructs {
	dbClient := models.InitDbClient()
	r := &Route{DbClient: dbClient}

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
