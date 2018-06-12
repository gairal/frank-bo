package routes

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/gairal/frank-gairal-bo/models"
)

// TravelsRoute Struct
type TravelsRoute struct{ *Route }

// index - Get all works
func (route *TravelsRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Travel{Entity: route.e}
	es := e.GetAll(appengine.NewContext(r))
	route.serveGetAll(w, &es)
}

// byKey - Get one entity by its key
func (route *TravelsRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Travel{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}

// add - Add one test entity
func (route *TravelsRoute) add(w http.ResponseWriter, r *http.Request) {
	e1 := models.Travel{
		Entity: route.e,
		Coordinates: appengine.GeoPoint{
			Lat: 0.66,
			Lng: 0.55,
		},
		Order: 0,
		Place: "paris",
	}

	ctx := appengine.NewContext(r)

	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "travel", nil), &e1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
