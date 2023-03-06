package main

import (
	"context"
	"flag"
	"fmt"

	dtrack "github.com/DependencyTrack/client-go"
)

func main() {
	var (
		apiKey   string
		authURL  string
		projName string
	)
	flag.StringVar(&apiKey, "apiKey", "", "dependency track api key")
	flag.StringVar(&authURL, "url", "", "dependency track base url")
	flag.StringVar(&projName, "name", "", "")

	flag.Parse()
	c, err := dtrack.NewClient(authURL, dtrack.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	r, err := c.Project.Create(context.Background(), dtrack.Project{
		Name:   projName,
		Active: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:%s\n", projName, r.UUID)
}
