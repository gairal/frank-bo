package models

import (
	"time"

	"cloud.google.com/go/datastore"
)

// Work - Work Structure
type Work struct {
	Entity          *Entity        `json:"-"`
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

// GetAll - Get All works
func (e *Work) GetAll() Works {
	q := datastore.NewQuery("work")
	var entities Works
	e.Entity.getAll(q, &entities, "date_in", true)

	img := &Image{Entity: e.Entity}
	imgs := img.getAllOrdered()

	for i := 0; i < len(entities); i++ {
		entities[i].Image = imgs[entities[i].ImageKey.ID]
	}

	return entities
}

// Get - Get Work by id
func (e *Work) Get(k int64) {
	e.Entity.get("work", k, e)
}
