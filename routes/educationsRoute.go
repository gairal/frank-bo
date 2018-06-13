package routes

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
	}
}
