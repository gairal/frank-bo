package models

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// Entity - abstract entity
type Entity struct {
	db *datastore.Client
}

// InitDbClient - Init Data Store Client
func InitDbClient() *datastore.Client {
	ctx := context.Background()
	projectID := "com-gairal-frank-api"

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		panic(err)
	}

	return client
}

// GetAll - Get All entities
func (e *Entity) GetAll(db *datastore.Client, kind string) Works {
	ctx := context.Background()

	q := datastore.NewQuery(kind)
	var works Works
	if _, err := db.GetAll(ctx, q, &works); err != nil {
		log.Fatalf("Failed to get works: %v", err)
		panic(err)
	}

	return works
}
