// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gairal/frank-gairal-bo/models"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/works", WorksIndex)

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// // WorksIndex index for works
// func WorksIndex(w http.ResponseWriter, r *http.Request) {
// 	todos := models.Works{
// 		models.Work{Name: "Write presentation"},
// 		models.Work{Name: "Host meetup"},
// 	}

// 	json.NewEncoder(w).Encode(todos)
// }

package main

import (
	"github.com/gairal/frank-gairal-bo/routes"
	"google.golang.org/appengine"
)

func main() {
	routes.RegisterHandlers()
	appengine.Main()

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
