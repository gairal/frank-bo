package main

import (
	"github.com/gairal/frank-gairal-bo/routes"
	"google.golang.org/appengine"
)

func main() {
	routes.RegisterHandlers()
	appengine.Main()
}
