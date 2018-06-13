package routes

// InterestsRoute Struct
type InterestsRoute struct{ *Route }

// index - Get all works
func (route *InterestsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"interestsIndex",
			"GET",
			"/interests",
			route.index,
		},
		RouteStruct{
			"interestsByKey",
			"GET",
			"/interests/{key}",
			route.byKey,
		},
	}
}
