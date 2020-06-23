package main

import (
	"log"

	"github.com/renishb10/golang-podcast/itunes"
)

func main() {
	ias := itunes.NewItunesApiService()

	res, err := ias.Search("Full Stack Radio")
	if err != nil {
		log.Fatalf("Error while Searching: %v", err)
	}

	for _, item := range res.Results {
		log.Println("-------------------------------")
		log.Printf("Artist: %s", item.ArtistName)
		log.Printf("Podcast Name: %s", item.TrackName)
		log.Printf("Feed URL: %s", item.FeedURL)
		log.Println("-------------------------------")
	}
}
