package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Category - category Structure
type Category struct {
	Key   *datastore.Key `json:"-"`
	Name  string         `json:"name" datastore:"name"`
	Order int            `json:"order" datastore:"order"`
}

// Categories - Array of Category
type Categories []Category

// OrderedCategories - Ordered map of Category
type OrderedCategories map[int64]Category

// GetAllCategories - Get All Categories
func (e *Category) GetAllCategories(ctx context.Context) Categories {
	q := datastore.NewQuery("category")
	var entities Categories
	keys := GetAll(ctx, q, &entities, "order", false)

	for i := 0; i < len(entities); i++ {
		entities[i].Key = keys[i]
	}

	return entities
}

// GetAll - Get All Categories
func (e *Category) GetAll(ctx context.Context) interface{} {
	return e.GetAllCategories(ctx)
}

// GetAllByCategory - Get All Skills by category
func (e *Category) GetAllByCategory(ctx context.Context) interface{} {
	return e.GetAll(ctx)
}

// getMulti - Get All Categories In a map whick keys are img keys
func (e *Category) getMulti(ctx context.Context, keys []*datastore.Key) Categories {
	var entities = make(Categories, len(keys))
	GetMulti(ctx, keys, entities)
	return entities
}

// getMultiOrdered - Get All Categories In a map whick keys are img keys
func (e *Category) getMultiOrdered(ctx context.Context, keys []*datastore.Key) OrderedCategories {
	imgs := e.getMulti(ctx, keys)

	orderedImgs := make(OrderedCategories)
	for i := 0; i < len(keys); i++ {
		orderedImgs[keys[i].IntID()] = imgs[i]
	}

	return orderedImgs
}

// Get - Get Education by id
func (e *Category) Get(ctx context.Context, k int64) {
	Get(ctx, "category", k, e)
}

// orderedEntitiesToSlice - Transforman education slice to IEntity slice
func (e *Category) orderedEntitiesToSlice(oi OrderedCategories) Categories {
	res := make(Categories, len(oi))
	idx := 0
	for _, v := range oi {
		img := v
		res[idx] = Category(img)
		idx++
	}

	return res
}
