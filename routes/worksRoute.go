package routes

// WorksRoute Struct
type WorksRoute struct{ *Route }

// index - Get all works
func (r *WorksRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"worksIndex",
			"GET",
			"/works",
			r.index,
		},
		RouteStruct{
			"worksByKey",
			"GET",
			"/works/{key}",
			r.byKey,
		},
	}
}
