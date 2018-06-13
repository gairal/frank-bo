package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Category - Catgory Structure
type Category struct {
	Name  string `json:"name" datastore:"name"`
	Order int    `json:"order" datastore:"order"`
}

// Categories - Array of Category
type Categories []Category

// OrderedCategories - Ordered map of Category
type OrderedCategories map[int64]Category

// GetAll - Get All Categories
func (e *Category) GetAll(ctx context.Context) []IEntity {
	q := datastore.NewQuery("image")
	var entities Categories
	GetAll(ctx, q, &entities, "", false)

	return e.sliceToIEntitySlice(entities)
}

// GetAllByCategory - Get All Skills by catgory
func (e *Category) GetAllByCategory(ctx context.Context) []IEntity {
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
	Get(ctx, "image", k, e)
}

// sliceToIEntitySlice - Transforman education slice to IEntity slice
func (e *Category) sliceToIEntitySlice(es Categories) []IEntity {
	res := make([]IEntity, len(es))

	for i, v := range es {
		ent := v
		res[i] = IEntity(&ent)
	}

	return res
}

// orderedCategoriesToIEntitySlice - Transforman education slice to IEntity slice
func (e *Category) orderedCategoriesToIEntitySlice(oi OrderedCategories) []IEntity {
	res := make([]IEntity, len(oi))
	idx := 0
	for _, v := range oi {
		img := v
		res[idx] = IEntity(&img)
		idx++
	}

	return res
}
