package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// StoriesTypes types of Stories list
var StoriesTypes = [...]string{
	"top",
	"new",
	"best",
	"ask",
	"show",
	"job",
}

const (
	// StoriesURL base URL for getting list of stories IDs by types, ie. "new", "top"...
	StoriesURL = "https://hacker-news.firebaseio.com/v0/%sstories.json"
	// ItemURL base URL for getting item detail
	ItemURL = "https://hacker-news.firebaseio.com/v0/item/%v.json"
)

// Item fields, see: https://github.com/HackerNews/API#items
type Item struct {
	ID          int
	Time        int
	Score       int
	Type        string
	By          string
	Text        string
	Dead        bool
	Parent      int
	Kids        []int
	URL         string
	Title       string
	Parts       interface{}
	Descendants interface{}
}

// GetStoryByID retreives a story's detail
func GetStoryByID(id int, target interface{}) error {
	url := fmt.Sprintf(ItemURL, id)
	return GetJSON(url, target)
}

// GetStoriesURL returns full URL based on base URL & stories types
func GetStoriesURL(stories string) (url string, err error) {
	for _, s := range StoriesTypes {
		if stories == s {
			url := fmt.Sprintf(StoriesURL, stories)
			return url, nil
		}
	}
	return "", fmt.Errorf("unknown stories type")
}

// GetJSON see: http://stackoverflow.com/a/31129967/4328963
func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
