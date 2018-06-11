package models

import (
	"cloud.google.com/go/datastore"
)

// Skill - Skill Structure
type Skill struct {
	Entity      *Entity        `json:"-"`
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

// GetAll - Get All Skills
func (e *Skill) GetAll() Skills {
	q := datastore.NewQuery("skill")
	var entities Skills
	e.Entity.getAll(q, &entities, "", false)

	return entities
}

// GetAllByCategory - Get All Skills by catgory
func (e *Skill) GetAllByCategory() SkillCategories {
	entities := e.GetAll()
	var entityCategories SkillCategories

	cat := &Category{Entity: e.Entity}
	cats := cat.getAll()

	entitiesMap := make(map[int64]Skills)

	// Create skills grouped by categories
	for i := 0; i < len(entities); i++ {
		if _, ok := entitiesMap[entities[i].CategoryKey.ID]; !ok {
			var skills Skills
			entitiesMap[entities[i].CategoryKey.ID] = skills
		}

		entitiesMap[entities[i].CategoryKey.ID] = append(entitiesMap[entities[i].CategoryKey.ID], entities[i])
	}

	for i := 0; i < len(cats); i++ {
		if val, ok := entitiesMap[cats[i].Key.ID]; ok {
			entityCategories = append(entityCategories, SkillCategory{cats[i], val})
		}
	}

	return entityCategories
}

// Get - Get Skill by id
func (e *Skill) Get(k int64) {
	e.Entity.get("skill", k, e)
}
