package routes

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
