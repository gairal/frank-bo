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

// InterestCategory - Category
type InterestCategory struct {
	Category
	Interests Interests `json:"interests" datastore:"interests"`
}

// InterestCategories - Array of InterestCategory
type InterestCategories []InterestCategory

// GetAllInterests - Get All Interests
func (e *Interest) GetAllInterests(ctx context.Context) Interests {
	q := datastore.NewQuery("interest")
	var entities Interests
	GetAll(ctx, q, &entities, "", false)

	return entities
}

// GetAll - Get All Interests
func (e *Interest) GetAll(ctx context.Context) interface{} {
	return e.GetAllInterests(ctx)
}

// Get - Get Interest by id
func (e *Interest) Get(ctx context.Context, k int64) {
	Get(ctx, "interest", k, e)
}

// GetAllByCategory - Get All Skills by category
func (e *Interest) GetAllByCategory(ctx context.Context) interface{} {
	entities := e.GetAllInterests(ctx)

	return e.GetCategories(ctx, entities)
}

// GetCategories - Get Categories in education slice
func (e *Interest) GetCategories(ctx context.Context, entities Interests) InterestCategories {
	// Get a map of all img keys in my educations
	cats := make(map[int64]*datastore.Key)

	for _, v := range entities {
		if _, ok := cats[v.CategoryKey.IntID()]; !ok {
			cats[v.CategoryKey.IntID()] = v.CategoryKey
		}
	}

	// // Get all Categories
	c := &Category{}
	orderedCats := c.GetAllCategories(ctx)

	// Create Interests grouped by categories
	entitiesMap := make(map[int64]Interests)
	for i := 0; i < len(entities); i++ {
		id := entities[i].CategoryKey.IntID()
		if _, ok := entitiesMap[id]; !ok {
			var interests Interests
			entitiesMap[id] = interests
		}

		entitiesMap[id] = append(entitiesMap[id], entities[i])
	}

	var entityCategories InterestCategories
	for i := 0; i < len(orderedCats); i++ {
		cat := orderedCats[i]
		key := cat.Key.IntID()
		if val, ok := entitiesMap[key]; ok {
			entityCategories = append(entityCategories, InterestCategory{cat, val})
		}
	}

	return entityCategories
}
