package routes

// WorksRoute Struct
type WorksRoute struct{ *Route }

// index - Get all works
func (route *WorksRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"worksIndex",
			"GET",
			"/works",
			route.index,
		},
		RouteStruct{
			"worksByKey",
			"GET",
			"/works/{key}",
			route.byKey,
		},
	}
}
