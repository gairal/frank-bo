package models

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// Entity - abstract entity
type Entity struct {
	DbClient *datastore.Client
}

// InitDbClient - Init Data Store Client
func InitDbClient() *Entity {
	ctx := context.Background()
	projectID := "com-gairal-frank-api"

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		panic(err)
	}

	e := &Entity{DbClient: client}

	return e
}

// GetAll - Return all of type entities
func (e *Entity) getAll(q *datastore.Query, entities interface{}) {
	ctx := context.Background()

	if _, err := e.DbClient.GetAll(ctx, q, entities); err != nil {
		log.Fatalf("Failed to get entities: %v", err)
		panic(err)
	}
}
