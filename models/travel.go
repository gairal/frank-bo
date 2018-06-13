package models

import (
	"context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Travel - Travel Structure
type Travel struct {
	Coordinates appengine.GeoPoint `json:"coordinates" datastore:"coordinates"`
	Order       int                `json:"order" datastore:"order"`
	Place       string             `json:"place" datastore:"place"`
}

// Travels - Array of Travel
type Travels []Travel

// GetAll - Get All Travels
func (e *Travel) GetAll(ctx context.Context) []IEntity {
	q := datastore.NewQuery("travel")
	var entities Travels
	GetAll(ctx, q, &entities, "order", false)

	res := make([]IEntity, len(entities))
	for i, v := range entities {
		res[i] = IEntity(&v)
	}

	return res
}

// func (e *Travel) GetAll(ctx context.Context) Travels {
// 	q := datastore.NewQuery("travel")
// 	var entities Travels
// 	GetAll(ctx, q, &entities, "", false)

// 	return entities
// }

// Get - Get Travel by id
func (e *Travel) Get(ctx context.Context, k int64) {
	Get(ctx, "travel", k, e)
}
