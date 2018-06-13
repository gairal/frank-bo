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
type Images []*Image

// OrderedImages - Ordered map of Image
type OrderedImages map[int64]Image

// GetAll - Get All Images
func (e *Image) GetAll(ctx context.Context) []IEntity {
	q := datastore.NewQuery("image")
	var entities Images
	GetAll(ctx, q, &entities, "", false)

	res := make([]IEntity, len(entities))
	for i, v := range entities {
		res[i] = IEntity(v)
	}

	return res
}

// GetAllByCategory - Get All Skills by catgory
func (e *Image) GetAllByCategory(ctx context.Context) []IEntity {
	return e.GetAll(ctx)
}

// getMulti - Get All Images In a map whick keys are img keys
func (e *Image) getMulti(ctx context.Context, keys []*datastore.Key) Images {
	var entities = make([]*Image, len(keys))
	GetMulti(ctx, keys, entities)
	return entities
}

// getMultiOrdered - Get All Images In a map whick keys are img keys
func (e *Image) getMultiOrdered(ctx context.Context, keys []*datastore.Key) OrderedImages {
	imgs := e.getMulti(ctx, keys)

	orderedImgs := make(OrderedImages)
	for i := 0; i < len(keys); i++ {
		orderedImgs[keys[i].IntID()] = *imgs[i]
	}

	return orderedImgs
}

// Get - Get Education by id
func (e *Image) Get(ctx context.Context, k int64) {
	Get(ctx, "image", k, e)
}

// orderedImagesToIEntitySlice - Transforman education slice to IEntity slice
func (e *Image) orderedImagesToIEntitySlice(oi OrderedImages) []IEntity {
	res := make([]IEntity, len(oi))
	idx := 0
	for _, v := range oi {
		img := v
		res[idx] = IEntity(&img)
		idx++
	}

	return res
}
