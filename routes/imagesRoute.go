package routes

// ImagesRoute Struct
type ImagesRoute struct{ *Route }

// index - Get all works
func (route *ImagesRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"imagesIndex",
			"GET",
			"/images",
			route.index,
		},
		RouteStruct{
			"imagesByKey",
			"GET",
			"/images/{key}",
			route.byKey,
		},
	}
}
