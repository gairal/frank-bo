package models

import (
	"context"
	"log"
	// "google.golang.org/appengine/datastore"
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
func (e *Entity) getAll(q *datastore.Query, entities interface{}, orderCol string, descending bool) {
	ctx := context.Background()

	if orderCol != "" {
		var order = ""
		if descending {
			order = "-"
		}

		q = q.Order(order + orderCol)
	}

	if _, err := e.DbClient.GetAll(ctx, q, entities); err != nil {
		log.Fatalf("Failed to get entities: %v", err)
		panic(err)
	}
}

// Get - Return an entity by its key
func (e *Entity) get(kind string, k int64, entity interface{}) {
	ctx := context.Background()

	key := datastore.IDKey(kind, k, nil)

	if err := e.DbClient.Get(ctx, key, entity); err != nil {
		log.Fatalf("Failed to get entity: %v", err)
		panic(err)
	}
}
