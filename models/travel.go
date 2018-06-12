package models

import (
	"context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Travel - Travel Structure
type Travel struct {
	Entity      *Entity            `json:"-" datastore:"-"`
	Coordinates appengine.GeoPoint `json:"coordinates" datastore:"coordinates"`
	Order       int                `json:"order" datastore:"order"`
	Place       string             `json:"place" datastore:"place"`
}

// Travels - Array of Travel
type Travels []Travel

// GetAll - Get All Travels
func (e *Travel) GetAll(ctx context.Context) Travels {
	q := datastore.NewQuery("travel")
	var entities Travels
	e.Entity.getAll(ctx, q, &entities, "", false)

	return entities
}

// Get - Get Travel by id
func (e *Travel) Get(ctx context.Context, k int64) {
	e.Entity.get(ctx, "travel", k, e)
}
