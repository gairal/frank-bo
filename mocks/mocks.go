package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gairal/frank-gairal-bo/models"
)

// getJSON - Get Raw Json
func getJSON(file string) []byte {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

func loadImages(ctx context.Context) {
	raw := getJSON("./images")
	var e models.Images
	json.Unmarshal(raw, &e)

	// TODO: Inject images
}

// InitDB - Inject fixture
func InitDB(ctx context.Context) {
	loadImages(ctx)
}
