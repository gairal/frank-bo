package routes

// TravelsRoute Struct
type TravelsRoute struct{ *Route }

// index - Get all works
func (r *TravelsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"travelsIndex",
			"GET",
			"/travels",
			r.index,
		},
		RouteStruct{
			"travelsByKey",
			"GET",
			"/travels/{key}",
			r.byKey,
		},
	}
}
