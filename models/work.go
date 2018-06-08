package models

import (
	"time"
)

// Work - Work Structure
type Work struct {
	Accomplishments string    `json:"accomplishments"`
	DateIn          time.Time `json:"date_in"`
	Description     time.Time `json:"description"`
	Image           string    `json:"image"`
	Name            string    `json:"name"`
	Order           uint8     `json:"order"`
	Place           string    `json:"place"`
	Title           string    `json:"title"`
	Website         string    `json:"website"`
}

// Works - Array or Work
type Works []Work
