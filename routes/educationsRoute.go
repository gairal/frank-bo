package routes

// EducationsRoute Struct
type EducationsRoute struct{ *Route }

// index - Get all works
func (r *EducationsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"educationsIndex",
			"GET",
			"/educations",
			r.index,
		},
		RouteStruct{
			"educationsByKey",
			"GET",
			"/educations/{key}",
			r.byKey,
		},
	}
}
