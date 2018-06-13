package routes

// TravelsRoute Struct
type TravelsRoute struct{ *Route }

// index - Get all works
func (route *TravelsRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"travelsIndex",
			"GET",
			"/travels",
			route.index,
		},
		RouteStruct{
			"travelsByKey",
			"GET",
			"/travels/{key}",
			route.byKey,
		},
	}
}
