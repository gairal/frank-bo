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

// GetAllSkills - Get All Skills
func (e *Skill) GetAllSkills(ctx context.Context) Skills {
	q := datastore.NewQuery("skill")
	var entities Skills
	GetAll(ctx, q, &entities, "", false)

	return entities
}

// GetAll - Get All Skills
func (e *Skill) GetAll(ctx context.Context) []IEntity {
	return e.sliceToIEntitySlice(e.GetAllSkills(ctx))
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

// GetAllByCategory - Get All Skills by catgory
func (e *Skill) GetAllByCategory(ctx context.Context) []IEntity {
	entities := e.GetAllSkills(ctx)

	// entities := make(Skills, reflect.ValueOf(ent).Len())
	// var entityCategories SkillCategories

	// cat := &Category{Entity: e.Entity}
	// cats := make(Categories, reflect.ValueOf(cat.GetAll()).Len())

	// entitiesMap := make(map[int64]Skills)

	// // Create skills grouped by categories
	// for i := 0; i < len(entities); i++ {
	// 	if _, ok := entitiesMap[entities[i].CategoryKey.ID]; !ok {
	// 		var skills Skills
	// 		entitiesMap[entities[i].CategoryKey.ID] = skills
	// 	}

	// 	entitiesMap[entities[i].CategoryKey.ID] = append(entitiesMap[entities[i].CategoryKey.ID], entities[i])
	// }

	// for i := 0; i < len(cats); i++ {
	// 	if val, ok := entitiesMap[cats[i].Key.ID]; ok {
	// 		entityCategories = append(entityCategories, SkillCategory{cats[i], val})
	// 	}
	// }

	// var res []IEntity
	// for i, v := range entities {
	// 	res[i] = IEntity(&v)
	// }

	return e.sliceToIEntitySlice(entities)
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
