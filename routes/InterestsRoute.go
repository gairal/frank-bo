package routes

// InterestsRoute Struct
type InterestsRoute struct{ *Route }

// index - Get all works
func (r *InterestsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"interestsIndex",
			"GET",
			"/interests",
			r.index,
		},
		RouteStruct{
			"interestsByKey",
			"GET",
			"/interests/{key}",
			r.byKey,
		},
	}
}
