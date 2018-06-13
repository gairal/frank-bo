package routes

// CategoriesRoute Struct
type CategoriesRoute struct{ *Route }

// index - Get all works
func (route *CategoriesRoute) getRoutes() RouteStructs {
	return RouteStructs{
		RouteStruct{
			"categoriesIndex",
			"GET",
			"/categories",
			route.index,
		},
		RouteStruct{
			"categoriesByKey",
			"GET",
			"/categories/{key}",
			route.byKey,
		},
	}
}
