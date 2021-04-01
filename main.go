package main

import (
	"fmt"
	"log"

	nyaa "github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	opt := nyaa.SearchOptions{}

	torrents, err := nyaa.SearchAll(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(torrents[0].Name)
	fmt.Println(torrents[1].Name)
	fmt.Println(torrents[2].Name)
}
