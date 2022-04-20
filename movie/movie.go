package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Only capitalized fields are marshalled
type Movie struct {
	Title  string
	Year   int  `json:"released"` // <- field tags
	Color  bool `json:"color,omitempty"`
	Actors []string
}

// field tags `json:"alternativeName"`

var movies = []Movie{
	{"Casablanca", 1942, false, []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{"Cool Hand Luke", 1967, true, []string{"Paul Newman"}},
	{"Bullitt", 1968, true, []string{"Steve McQueen", "Jacqualine Bisset"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed %s", err)
	}
	fmt.Printf("%s\n", data)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshalling failed %s", err)
	}
	fmt.Println(titles)
}
