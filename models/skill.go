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

// GetAllByCategory - Get All Skills by category
func (e *Skill) GetAllByCategory(ctx context.Context) interface{} {
	entities := e.GetAllSkills(ctx)

	return e.GetCategories(ctx, entities)
}

// GetCategories - Get Categories in education slice
func (e *Skill) GetCategories(ctx context.Context, entities Skills) SkillCategories {
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

	// Create skills grouped by categories
	entitiesMap := make(map[int64]Skills)
	for i := 0; i < len(entities); i++ {
		id := entities[i].CategoryKey.IntID()
		if _, ok := entitiesMap[id]; !ok {
			var skills Skills
			entitiesMap[id] = skills
		}

		entitiesMap[id] = append(entitiesMap[id], entities[i])
	}

	var entityCategories SkillCategories
	for i := 0; i < len(orderedCats); i++ {
		cat := orderedCats[i]
		key := cat.Key.IntID()
		if val, ok := entitiesMap[key]; ok {
			entityCategories = append(entityCategories, SkillCategory{cat, val})
		}
	}

	return entityCategories
}
