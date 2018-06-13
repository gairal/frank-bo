package models

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// Skill - Skill Structure
type Skill struct {
	CategoryKey *datastore.Key `json:"-" datastore:"category"`
	Description string         `json:"description" datastore:"description"`
	Level       int            `json:"level" datastore:"level"`
	Name        string         `json:"name" datastore:"name"`
}

// Skills - Array of Skill
type Skills []Skill

// GetAll - Get All Skills
func (e *Skill) GetAll(ctx context.Context) []IEntity {
	q := datastore.NewQuery("skill")
	var entities Skills
	GetAll(ctx, q, &entities, "", false)

	return e.sliceToIEntitySlice(entities)
}

// sliceToIEntitySlice - Transforman education slice to IEntity slice
func (e *Skill) sliceToIEntitySlice(es Skills) []IEntity {
	res := make([]IEntity, len(es))

	for i, v := range es {
		ent := v
		res[i] = IEntity(&ent)
	}

	return res
}

// Get - Get Skill by id
func (e *Skill) Get(ctx context.Context, k int64) {
	Get(ctx, "skill", k, e)
}
