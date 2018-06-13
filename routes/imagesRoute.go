package routes

// ImagesRoute Struct
type ImagesRoute struct{ *Route }

// index - Get all works
func (r *ImagesRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"imagesIndex",
			"GET",
			"/images",
			r.index,
		},
		RouteStruct{
			"imagesByKey",
			"GET",
			"/images/{key}",
			r.byKey,
		},
	}
}
