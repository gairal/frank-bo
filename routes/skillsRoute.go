package routes

// SkillsRoute Struct
type SkillsRoute struct{ *Route }

// index - Get all works
func (route *SkillsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"skillsIndex",
			"GET",
			"/skills",
			route.index,
		},
		RouteStruct{
			"skillsByKey",
			"GET",
			"/skills/{key}",
			route.byKey,
		},
	}
}
