package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	opt := nyaa.SearchOptions{
		Provider: "nyaa", // Provider is the only required option
		Query:    "jojo",
		Category: "anime",
		SortBy:   "downloads",
		Filter:   "trusted-only",
	}

	torrents, err := nyaa.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range torrents {
		if i > 4 { //print first 5
			break
		}
		fmt.Printf("%+v", v)
	}
}
