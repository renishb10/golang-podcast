// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type FeedItem struct {
	PubDate     string  `json:"pubDate"`
	Text        string  `json:"text"`
	Title       string  `json:"title"`
	SubTitle    string  `json:"subTitle"`
	Description string  `json:"description"`
	Image       *string `json:"image"`
	Summary     string  `json:"summary"`
	Link        string  `json:"link"`
	Duration    string  `json:"duration"`
}

type Podcast struct {
	Artist        string   `json:"artist"`
	PodcastName   string   `json:"podcastName"`
	FeedURL       string   `json:"feedUrl"`
	Thumbnail     string   `json:"thumbnail"`
	EpisodesCount int      `json:"episodesCount"`
	Genres        []string `json:"genres"`
}
