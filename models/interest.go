package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Interest - Interest Structure
type Interest struct {
	CategoryKey *datastore.Key `json:"-" datastore:"category"`
	Description string         `json:"description" datastore:"description"`
	Order       int            `json:"order" datastore:"order"`
	Name        string         `json:"name" datastore:"name"`
}

// Interests - Array of Interest
type Interests []Interest

// GetAll - Get All Interests
func (e *Interest) GetAll(ctx context.Context) interface{} {
	q := datastore.NewQuery("interest")
	var entities Interests
	GetAll(ctx, q, &entities, "order", false)

	return entities
}

// GetAllByCategory - Get All Skills by catgory
func (e *Interest) GetAllByCategory(ctx context.Context) interface{} {
	return e.GetAll(ctx)
}

// Get - Get Interest by id
func (e *Interest) Get(ctx context.Context, k int64) {
	Get(ctx, "interest", k, e)
}
