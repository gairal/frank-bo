package models

import (
	"time"

	"cloud.google.com/go/datastore"
)

// Education - Education Structure
type Education struct {
	Entity           *Entity        `json:"-"`
	ShortDescription string         `json:"short_description" datastore:"short_description"`
	YearIn           time.Time      `json:"year_in" datastore:"year_in"`
	YearOut          time.Time      `json:"year_out" datastore:"year_out"`
	Diploma          string         `json:"diploma" datastore:"diploma"`
	Place            string         `json:"place" datastore:"place"`
	Image            *datastore.Key `json:"image" datastore:"image"`
	Name             string         `json:"name" datastore:"name"`
	Website          string         `json:"website" datastore:"website"`
	Extra            string         `json:"extra" datastore:"extra"`
}

// Educations - Array of Education
type Educations []Education

// GetAll - Get All Educations
func (e *Education) GetAll() interface{} {
	q := datastore.NewQuery("education")
	var entities Educations
	e.Entity.getAll(q, &entities, "year_in", true)
	return entities
}
