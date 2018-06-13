package routes

// CategoriesRoute Struct
type CategoriesRoute struct{ *Route }

// index - Get all works
func (r *CategoriesRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"categoriesIndex",
			"GET",
			"/categories",
			r.index,
		},
		RouteStruct{
			"categoriesByKey",
			"GET",
			"/categories/{key}",
			r.byKey,
		},
	}
}
