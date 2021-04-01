package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	description, err := nyaa.TorrentDescription(1229657, "nyaa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(description)
}
