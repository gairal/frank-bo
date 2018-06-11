package models

import "cloud.google.com/go/datastore"

// Image - Education Structure
type Image struct {
	Entity *Entity        `json:"-"`
	Key    *datastore.Key `json:"-" datastore:"__key__"`
	Name   string         `json:"name" datastore:"name"`
}

// Images - Array of Image
type Images []Image

// OrderedImages - Ordered map of Image
type OrderedImages map[int64]Image

// getAll - Get All Images
func (e *Image) getAll() Images {
	q := datastore.NewQuery("image")
	var entities Images
	e.Entity.getAll(q, &entities, "", false)
	return entities
}

// getAllOrdered - Get All Images In a map whick keys are img keys
func (e *Image) getAllOrdered() OrderedImages {
	imgs := e.getAll()

	orderedImgs := make(map[int64]Image)
	for i := 0; i < len(imgs); i++ {
		orderedImgs[imgs[i].Key.ID] = imgs[i]
	}

	return orderedImgs
}
