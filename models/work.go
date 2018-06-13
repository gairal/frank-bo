package models

import (
	"context"
	"time"

	"google.golang.org/appengine/datastore"
)

// Work - Work Structure
type Work struct {
	Accomplishments string         `json:"accomplishments" datastore:"accomplishments"`
	DateIn          time.Time      `json:"date_in" datastore:"date_in"`
	DateOut         time.Time      `json:"date_out" datastore:"date_out"`
	Description     string         `json:"description" datastore:"description"`
	ImageKey        *datastore.Key `json:"-" datastore:"image"`
	Image           Image          `json:"image"`
	Name            string         `json:"name" datastore:"name"`
	Order           int            `json:"order" datastore:"order"`
	Place           string         `json:"place" datastore:"place"`
	Title           string         `json:"title" datastore:"title"`
	Website         string         `json:"website" datastore:"website"`
}

// Works - Array of Work
type Works []Work

// GetAll - Get All Works
func (e *Work) GetAll(ctx context.Context) []IEntity {
	q := datastore.NewQuery("work")

	// Get All Works
	var entities Works
	GetAll(ctx, q, &entities, "", false)

	e.GetImages(ctx, entities)

	return e.sliceToIEntitySlice(entities)
}

// GetAllByCategory - Get All Skills by catgory
func (e *Work) GetAllByCategory(ctx context.Context) []IEntity {
	return e.GetAll(ctx)
}

// Get - Get work by id
func (e *Work) Get(ctx context.Context, k int64) {
	Get(ctx, "work", k, e)
}

// GetImages - Get Images in work slice
func (e *Work) GetImages(ctx context.Context, entities Works) {
	// Get a map of all img keys in my Works
	imgs := make(map[int64]*datastore.Key)

	for _, v := range entities {
		if _, ok := imgs[v.ImageKey.IntID()]; !ok {
			imgs[v.ImageKey.IntID()] = v.ImageKey
		}
	}

	// Get all Images
	orderedImgs := GetOrderedImgs(ctx, imgs)

	// Set Image to each work
	for i := 0; i < len(entities); i++ {
		entities[i].Image = orderedImgs[entities[i].ImageKey.IntID()]
	}
}

// sliceToIEntitySlice - Transforman work slice to IEntity slice
func (e *Work) sliceToIEntitySlice(es Works) []IEntity {
	res := make([]IEntity, len(es))

	for i, v := range es {
		ent := v
		res[i] = IEntity(&ent)
	}

	return res
}
