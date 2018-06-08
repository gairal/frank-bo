package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// WorksIndex - Get all works
func WorksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ctx := appengine.NewContext(r)

	// q := datastore.NewQuery("work")
	// var works []models.Work
	// if _, err := q.GetAll(ctx, &works); err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	panic(err)
	// 	// Handle error.
	// }

	works := models.Works{
		models.Work{Name: "Write presentation"},
		models.Work{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(works)
}
