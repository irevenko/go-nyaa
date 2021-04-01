package main

import (
	"fmt"
	"log"

	nyaa "github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	opt := nyaa.SearchOptions{
		Query:    "haikyuu",
		Category: "live-action",
		SortBy:   "seeders",
	}

	torrents, err := nyaa.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(torrents[0])
}
