package models

import "cloud.google.com/go/datastore"

// Category - Category Structure
type Category struct {
	Entity *Entity        `json:"-"`
	Key    *datastore.Key `json:"-" datastore:"__key__"`
	Name   string         `json:"name" datastore:"name"`
	Order  int            `json:"order" datastore:"order"`
}

// Categories - Array of Category
type Categories []Category

// OrderedCategories - Ordered map of Category
type OrderedCategories map[int64]Category

// getAll - Get All Categories
func (e *Category) getAll() Categories {
	q := datastore.NewQuery("category")
	var entities Categories
	e.Entity.getAll(q, &entities, "order", false)
	return entities
}

// getAllOrdered - Get All Categories In a map whick keys are img keys
func (e *Category) getAllOrdered() OrderedCategories {
	cats := e.getAll()

	var orderedCats OrderedCategories
	orderedCats = make(map[int64]Category)
	for i := 0; i < len(cats); i++ {
		orderedCats[cats[i].Key.ID] = cats[i]
	}

	return orderedCats
}
