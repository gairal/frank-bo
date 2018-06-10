package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// Route - route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - all routes
type Routes []Route

// GetRoutes - List all routes
func GetRoutes() Routes {
	db := models.InitDbClient()
	works := &WorksRoute{db: db}

	return Routes{
		Route{
			"WorksIndex",
			"GET",
			"/works",
			works.worksIndex,
		},
	}
}
