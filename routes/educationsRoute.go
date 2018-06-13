package routes

import (
	"net/http"
	"time"

	"github.com/gairal/frank-gairal-bo/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// EducationsRoute Struct
type EducationsRoute struct{ *Route }

// index - Get all works
func (route *EducationsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"educationsIndex",
			"GET",
			"/educations",
			route.index,
		},
		RouteStruct{
			"educationsByKey",
			"GET",
			"/educations/{key}",
			route.byKey,
		},
		RouteStruct{
			"educationsAdd",
			"POST",
			"/educations",
			route.add,
		},
	}
}

// add - Add one test entity
func (route *EducationsRoute) add(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	k1 := datastore.NewKey(ctx, "image", "", 0, nil)
	i1 := models.Image{
		Name: "ekino",
	}

	ki1, err := datastore.Put(ctx, k1, &i1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	k2 := datastore.NewKey(ctx, "image", "", 0, nil)
	i2 := models.Image{
		Name: "backelite",
	}

	ki2, err := datastore.Put(ctx, k2, &i2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e1 := models.Education{
		ShortDescription: "Foo",
		YearIn:           time.Now(),
		YearOut:          time.Now(),
		Diploma:          "osef",
		Place:            "toto",
		Name:             "toto",
		Website:          "tata",
		ImageKey:         ki1,
	}
	e2 := models.Education{
		ShortDescription: "Bar",
		YearIn:           time.Now(),
		YearOut:          time.Now(),
		Diploma:          "osef",
		Place:            "tata",
		Name:             "toto",
		Website:          "tata",
		ImageKey:         ki2,
	}

	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "education", nil), &e1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "education", nil), &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
