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

// SkillCategory - Category
type SkillCategory struct {
	Category
	Skills Skills `json:"skills" datastore:"skills"`
}

// SkillCategories - Array of SkillCategory
type SkillCategories []SkillCategory

// GetAllSkills - Get All Skills
func (e *Skill) GetAllSkills(ctx context.Context) Skills {
	q := datastore.NewQuery("skill")
	var entities Skills
	GetAll(ctx, q, &entities, "", false)

	return entities
}

// GetAll - Get All Skills
func (e *Skill) GetAll(ctx context.Context) interface{} {
	return e.GetAllSkills(ctx)
}

// Get - Get Skill by id
func (e *Skill) Get(ctx context.Context, k int64) {
	Get(ctx, "skill", k, e)
}

// GetAllByCategory - Get All Skills by catgory
func (e *Skill) GetAllByCategory(ctx context.Context) interface{} {
	entities := e.GetAllSkills(ctx)

	cats := make(map[int64]*datastore.Key)

	for _, v := range entities {
		if _, ok := cats[v.CategoryKey.IntID()]; !ok {
			cats[v.CategoryKey.IntID()] = v.CategoryKey
		}
	}

	// log.Fatalf("Failed to get route")

	// Get all Categories
	orderedCats := GetOrderedCats(ctx, cats)
	cat := &Category{}
	return cat.orderedEntitiesToSlice(orderedCats)

	// return e.sliceToIEntitySlice(entities)
}

// GetCategories - Get Categories in education slice
func (e *Skill) GetCategories(ctx context.Context, entities Skills) {
	// Get a map of all img keys in my educations
	cats := make(map[int64]*datastore.Key)

	for _, v := range entities {
		if _, ok := cats[v.CategoryKey.IntID()]; !ok {
			cats[v.CategoryKey.IntID()] = v.CategoryKey
		}
	}

	// Get all Categories
	// orderedCats := GetOrderedCats(ctx, cats)

	// Set Category to each edu
	// for i := 0; i < len(entities); i++ {
	// 	entities[i].Category = orderedCats[entities[i].CategoryKey.IntID()]
	// }
}
