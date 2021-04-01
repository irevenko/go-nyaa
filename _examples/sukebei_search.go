package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	opt := nyaa.SearchOptions{
		Provider: "sukebei", // Provider is the only required option
		Category: "art-doujinshi",
		SortBy:   "comments",
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
