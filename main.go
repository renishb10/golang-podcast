package main

import (
	"log"

	"github.com/renishb10/golang-podcast/feeds"
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

		feed, err := feeds.GetFeed(item.FeedURL)
		if err != nil {
			log.Fatalf("\tError fetching Feeds: %v", err)
		}

		for _, pod := range feed.Channel.Item {
			log.Println("\t----------------------------")
			log.Printf("\tTitle: %s", pod.Title)
			log.Printf("\tDuration: %s", pod.Duration)
			log.Printf("\tDescription: %s", pod.Description)
			log.Printf("\tURL: %s", pod.Enclosure.URL)
			log.Println("\t----------------------------")
		}

		log.Println("-------------------------------")
	}
}
