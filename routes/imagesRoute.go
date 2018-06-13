package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// ImagesRoute Struct
type ImagesRoute struct{ *Route }

// index - Get all works
func (route *ImagesRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"imagesIndex",
			"GET",
			"/images",
			route.index,
		},
		RouteStruct{
			"imagesByKey",
			"GET",
			"/images/{key}",
			route.byKey,
		},
		RouteStruct{
			"imagesAdd",
			"POST",
			"/images",
			route.add,
		},
	}
}

// // add - Add one test entity
func (route *ImagesRoute) add(w http.ResponseWriter, r *http.Request) {
	e1 := models.Image{
		Name: "ekino",
	}

	ctx := appengine.NewContext(r)

	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "image", nil), &e1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
