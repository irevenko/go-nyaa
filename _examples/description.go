package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	description, err := nyaa.TorrentDescription("https://nyaa.si/view/1229657")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(description)
}
