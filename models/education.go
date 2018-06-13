package models

import (
	"context"
	"time"

	"google.golang.org/appengine/datastore"
)

// Education - Education Structure
type Education struct {
	ShortDescription string         `json:"short_description" datastore:"short_description"`
	YearIn           time.Time      `json:"year_in" datastore:"year_in"`
	YearOut          time.Time      `json:"year_out" datastore:"year_out"`
	Diploma          string         `json:"diploma" datastore:"diploma"`
	Place            string         `json:"place" datastore:"place"`
	ImageKey         *datastore.Key `json:"-" datastore:"image"`
	Image            Image          `json:"image"`
	Name             string         `json:"name" datastore:"name"`
	Website          string         `json:"website" datastore:"website"`
	Extra            string         `json:"extra" datastore:"extra"`
}

// Educations - Array of Education
type Educations []Education

// GetAll - Get All Educations
func (e *Education) GetAll(ctx context.Context) interface{} {
	q := datastore.NewQuery("education")

	// Get All educations
	var entities Educations
	GetAll(ctx, q, &entities, "year_in", true)

	e.GetImages(ctx, entities)

	return entities
}

// GetAllByCategory - Get All Skills by catgory
func (e *Education) GetAllByCategory(ctx context.Context) interface{} {
	return e.GetAll(ctx)
}

// Get - Get Education by id
func (e *Education) Get(ctx context.Context, k int64) {
	Get(ctx, "education", k, e)
}

// GetImages - Get Images in education slice
func (e *Education) GetImages(ctx context.Context, entities Educations) {
	// Get a map of all img keys in my educations
	imgs := make(map[int64]*datastore.Key)

	for _, v := range entities {
		if _, ok := imgs[v.ImageKey.IntID()]; !ok {
			imgs[v.ImageKey.IntID()] = v.ImageKey
		}
	}

	// Get all Images
	orderedImgs := GetOrderedImgs(ctx, imgs)

	// Set Image to each edu
	for i := 0; i < len(entities); i++ {
		entities[i].Image = orderedImgs[entities[i].ImageKey.IntID()]
	}
}
