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
	Image           *datastore.Key `json:"image" datastore:"image"`
	Name            string         `json:"name" datastore:"name"`
	Order           int            `json:"order" datastore:"order"`
	Place           string         `json:"place" datastore:"place"`
	Title           string         `json:"title" datastore:"title"`
	Website         string         `json:"website" datastore:"website"`
}

// Works - Array or Work
type Works []Work

// GetAll - Get All works
func (w *Work) GetAll() interface{} {
	q := datastore.NewQuery("work")
	var entities Works
	w.Entity.getAll(q, &entities, "date_in", true)
	return entities
}
