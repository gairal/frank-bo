package models

import (
	"cloud.google.com/go/datastore"
)

// Coordinates - Coordinates Structure
type Coordinates struct {
	Latitude  float32 `json:"latitude" datastore:"latitude"`
	Longitude float32 `json:"longitude" datastore:"longitude"`
}

// Travel - Travel Structure
type Travel struct {
	Entity      *Entity      `json:"-"`
	Coordinates *Coordinates `json:"Coordinates" datastore:"coordinates"`
	Order       int          `json:"order" datastore:"order"`
	Place       string       `json:"place" datastore:"place"`
}

// Travels - Array of Travel
type Travels []Travel

// GetAll - Get All Travels
func (e *Travel) GetAll() Travels {
	q := datastore.NewQuery("travel")
	var entities Travels
	e.Entity.getAll(q, &entities, "order", false)

	return entities
}

// Get - Get Travel by id
func (e *Travel) Get(k int64) {
	e.Entity.get("travel", k, e)
}
