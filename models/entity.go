package models

import (
	"context"
	"log"

	"google.golang.org/appengine/datastore"
)

// IEntity - entity interface
type IEntity interface {
	GetAll(context.Context) []IEntity
	Get(context.Context, int64)
}

// Keys - array of datastore keys
type Keys []*datastore.Key

// GetAll - Return all of type entities
func GetAll(ctx context.Context, q *datastore.Query, entities interface{}, orderCol string, descending bool) []*datastore.Key {
	// ctx := context.Background()

	if orderCol != "" {
		var order = ""
		if descending {
			order = "-"
		}

		q = q.Order(order + orderCol)
	}

	keys, err := q.GetAll(ctx, entities)

	if err != nil {
		log.Fatalf("Failed to get entities: %v", err)
		panic(err)
	}

	return keys
}

// Get - Return an entity by its key
func Get(ctx context.Context, kind string, k int64, entity interface{}) {
	key := datastore.NewKey(ctx, kind, "", k, nil)

	if err := datastore.Get(ctx, key, entity); err != nil {
		log.Fatalf("Failed to get entity: %v", err)
		panic(err)
	}
}

// GetMulti - Return multiple entiies based on their key
func GetMulti(ctx context.Context, keys []*datastore.Key, entities interface{}) {
	if err := datastore.GetMulti(ctx, keys, entities); err != nil {
		log.Fatalf("Failed to get multi entities: %v", err)
		panic(err)
	}
}

// GetOrderedImgs - Get Images in a map for an []IEntity
func GetOrderedImgs(ctx context.Context, keys map[int64]*datastore.Key) OrderedImages {
	imgsSlice := make(Keys, len(keys))
	idx := 0
	for _, v := range keys {
		imgsSlice[idx] = v
		idx++
	}

	img := &Image{}
	return img.getMultiOrdered(ctx, imgsSlice)
}

// GetOrderedCats - Get Images in a map for an []IEntity
func GetOrderedCats(ctx context.Context, keys map[int64]*datastore.Key) OrderedCategories {
	catsSlice := make(Keys, len(keys))
	idx := 0
	for _, v := range keys {
		catsSlice[idx] = v
		idx++
	}

	cat := &Category{}
	return cat.getMultiOrdered(ctx, catsSlice)
}
