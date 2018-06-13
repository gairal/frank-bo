package routes

import (
	"net/http"

	"google.golang.org/appengine"
)

// SkillsRoute Struct
type SkillsRoute struct{ *Route }

// index - Get all works
func (r *SkillsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"skillsIndex",
			"GET",
			"/skills",
			r.index,
		},
		RouteStruct{
			"skillsByCategory",
			"GET",
			"/skills/categories",
			r.byCategory,
		},
		RouteStruct{
			"skillsByKey",
			"GET",
			"/skills/{key}",
			r.byKey,
		},
	}
}

// byCategory - Get one entity by its key
func (r *SkillsRoute) byCategory(w http.ResponseWriter, req *http.Request) {
	es := r.entity.GetAllByCategory(appengine.NewContext(req))
	r.router.serve(w, &es)
}
