package models

import (
	"cloud.google.com/go/datastore"
)

// Interest - Interest Structure
type Interest struct {
	Entity      *Entity        `json:"-"`
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

// GetAll - Get All Interests
func (e *Interest) GetAll() Interests {
	q := datastore.NewQuery("interest")
	var entities Interests
	e.Entity.getAll(q, &entities, "order", false)

	return entities
}

// GetAllByCategory - Get All Interests by category
func (e *Interest) GetAllByCategory() InterestCategories {
	entities := e.GetAll()
	var entityCategories InterestCategories

	cat := &Category{Entity: e.Entity}
	cats := cat.getAll()

	entitiesMap := make(map[int64]Interests)

	// Create Interests grouped by categories
	for i := 0; i < len(entities); i++ {
		if _, ok := entitiesMap[entities[i].CategoryKey.ID]; !ok {
			var interests Interests
			entitiesMap[entities[i].CategoryKey.ID] = interests
		}

		entitiesMap[entities[i].CategoryKey.ID] = append(entitiesMap[entities[i].CategoryKey.ID], entities[i])
	}

	for i := 0; i < len(cats); i++ {
		if val, ok := entitiesMap[cats[i].Key.ID]; ok {
			entityCategories = append(entityCategories, InterestCategory{cats[i], val})
		}
	}

	return entityCategories
}

// Get - Get Interest by id
func (e *Interest) Get(k int64) {
	e.Entity.get("interest", k, e)
}
