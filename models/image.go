package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Image - Image Structure
type Image struct {
	// Key  *datastore.Key `json:"key" datastore:"__key__"`
	Name string `json:"name" datastore:"name"`
}

// Images - Array of Image
type Images []Image

// OrderedImages - Ordered map of Image
type OrderedImages map[int64]Image

// GetAll - Get All Images
func (e *Image) GetAll(ctx context.Context) interface{} {
	q := datastore.NewQuery("image")
	var entities Images
	GetAll(ctx, q, &entities, "", false)

	return entities
}

// GetAllByCategory - Get All Skills by category
func (e *Image) GetAllByCategory(ctx context.Context) interface{} {
	return e.GetAll(ctx)
}

// getMulti - Get All Images In a map whick keys are img keys
func (e *Image) getMulti(ctx context.Context, keys []*datastore.Key) Images {
	var entities = make([]Image, len(keys))
	GetMulti(ctx, keys, entities)
	return entities
}

// getMultiOrdered - Get All Images In a map whick keys are img keys
func (e *Image) getMultiOrdered(ctx context.Context, keys []*datastore.Key) OrderedImages {
	imgs := e.getMulti(ctx, keys)

	orderedImgs := make(OrderedImages)
	for i := 0; i < len(keys); i++ {
		orderedImgs[keys[i].IntID()] = imgs[i]
	}

	return orderedImgs
}

// Get - Get Education by id
func (e *Image) Get(ctx context.Context, k int64) {
	Get(ctx, "image", k, e)
}

// orderedEntitiesToSlice - Transforman education slice to IEntity slice
func (e *Image) orderedEntitiesToSlice(oi OrderedImages) Images {
	res := make(Images, len(oi))
	idx := 0
	for _, v := range oi {
		img := v
		res[idx] = Image(img)
		idx++
	}

	return res
}
