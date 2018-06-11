package routes

import (
	"net/http"

	"github.com/gairal/frank-gairal-bo/models"
)

// SkillsRoute Struct
type SkillsRoute struct{ *Route }

// index - Get all skills
func (route *SkillsRoute) index(w http.ResponseWriter, r *http.Request) {
	e := &models.Skill{Entity: route.e}
	es := e.GetAll()
	route.serveGetAll(w, es)
}

// byCategory - Get all skills by categories
func (route *SkillsRoute) byCategory(w http.ResponseWriter, r *http.Request) {
	e := &models.Skill{Entity: route.e}
	es := e.GetAllByCategory()
	route.serveGetAll(w, es)
}

// byKey - Get one entity by its key
func (route *SkillsRoute) byKey(w http.ResponseWriter, r *http.Request) {
	e := models.Skill{Entity: route.e}
	e.Get(route.getKey(r))
	route.serveGetAll(w, e)
}
