// TODO:
// - check other similar projects to see how they structure their go projects
// - further enchance this program ? - save retreived items in a map, maybe
// - try making a microservice with an ORM (like 'gorm' project)

package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gohn/stories"
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

func main() {
	fmt.Println("HN Reader")
	args := os.Args[1:]
	fmt.Printf("--- %s ---\n", args)

	// Find the right URL based on user input of stories types, ie. "new", "top"...
	url, err := stories.GetStoriesURL(args[0])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Get the list of stories
	var itemIDs []int
	if err := stories.GetJSON(url, &itemIDs); err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Get details of all stories in the list
	var items = make([]Item, len(itemIDs))
	var wg sync.WaitGroup
	wg.Add(len(items))
	for i, e := range itemIDs {
		// Due to 'go' keyword that run the func() in the background,
		// it may run when the for loop reached the max index.
		// Anonymous functions have access to outer scope (for loop) variables,
		// but 'when' they access them is unpredictable.
		// see: http://oyvindsk.com/writing/common-golang-mistakes-1
		ii, ee := i, e
		go func() {
			defer wg.Done()
			if err := stories.GetStoryByID(ee, &items[ii]); err != nil {
				log.Fatalf("Error: %s", err)
			}
		}()
	}
	wg.Wait()

	PrintItems(items)
}

// PrintItems prints items index & title
func PrintItems(items []Item) {
	for i, e := range items {
		fmt.Printf("[%v] %s\n", i+1, e.Title)
	}
}
