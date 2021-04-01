package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	files, err := nyaa.TorrentFiles(932413, "nyaa")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		fmt.Println(v)
	}
}
