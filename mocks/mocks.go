package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gairal/frank-gairal-bo/models"
	"google.golang.org/appengine/datastore"
)

// Mock - Mock class
type Mock struct{}

// getJSON - Get Raw Json
func (m *Mock) getJSON(file string) []byte {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

func (m *Mock) loadImages(ctx context.Context) {
	raw := m.getJSON("./images")
	var e models.Images
	json.Unmarshal(raw, &e)

	// TODO: Inject images
}

func (m *Mock) loadEducations(ctx context.Context) {
	ki1 := datastore.NewKey(ctx, "image", "", 0, nil)
	ki2 := datastore.NewKey(ctx, "image", "", 0, nil)
	i1 := models.Image{
		Name: "ekino",
	}

	i2 := models.Image{
		Name: "backelite",
	}

	keys, err := datastore.PutMulti(ctx, []*datastore.Key{ki1, ki2}, []*models.Image{&i1, &i2})
	if err != nil {
		log.Fatalf("Failed to put entity: %v", err)
		return
	}

	e1 := models.Education{
		ShortDescription: "Foo",
		YearIn:           time.Now(),
		YearOut:          time.Now(),
		Diploma:          "osef",
		Place:            "toto",
		Name:             "toto",
		Website:          "tata",
		ImageKey:         keys[0],
	}
	e2 := models.Education{
		ShortDescription: "Bar",
		YearIn:           time.Now(),
		YearOut:          time.Now(),
		Diploma:          "osef",
		Place:            "tata",
		Name:             "toto",
		Website:          "tata",
		ImageKey:         keys[1],
	}

	ke1 := datastore.NewIncompleteKey(ctx, "education", nil)
	ke2 := datastore.NewIncompleteKey(ctx, "education", nil)

	if _, err := datastore.PutMulti(ctx, []*datastore.Key{ke1, ke2}, []*models.Education{&e1, &e2}); err != nil {
		log.Fatalf("Failed to put entity: %v", err)
		return
	}
}

func (m *Mock) loadSkills(ctx context.Context) {
	kc1 := datastore.NewKey(ctx, "category", "", 0, nil)
	kc2 := datastore.NewKey(ctx, "category", "", 0, nil)

	c1 := models.Category{
		Name:  "ekino",
		Order: 0,
	}

	c2 := models.Category{
		Name:  "backelite",
		Order: 1,
	}

	keys, err := datastore.PutMulti(ctx, []*datastore.Key{kc1, kc2}, []*models.Category{&c1, &c2})
	if err != nil {
		log.Fatalf("Failed to put entity: %v", err)
		return
	}

	s1 := models.Skill{
		CategoryKey: keys[0],
		Name:        "toto",
	}
	s2 := models.Skill{
		CategoryKey: keys[0],
		Name:        "tata",
	}

	s3 := models.Skill{
		CategoryKey: keys[1],
		Name:        "titi",
	}
	s4 := models.Skill{
		CategoryKey: keys[1],
		Name:        "tutu",
	}

	ks1 := datastore.NewIncompleteKey(ctx, "skill", nil)
	ks2 := datastore.NewIncompleteKey(ctx, "skill", nil)
	ks3 := datastore.NewIncompleteKey(ctx, "skill", nil)
	ks4 := datastore.NewIncompleteKey(ctx, "skill", nil)

	if _, err := datastore.PutMulti(ctx, []*datastore.Key{ks1, ks2, ks3, ks4}, []*models.Skill{&s1, &s2, &s3, &s4}); err != nil {
		log.Fatalf("Failed to put entity: %v", err)
		return
	}
}

// InitDB - Inject fixture
func (m *Mock) InitDB(ctx context.Context) {
	// m.loadImages(ctx)
	// m.loadEducations(ctx)
	m.loadSkills(ctx)
}
