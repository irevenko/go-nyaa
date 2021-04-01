package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	f, err := nyaa.TorrentFiles(1366007, "nyaa")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f)

	// opt := nyaa.SearchOptions{
	// 	Provider: "nyaa",
	// }

	// torrents, err := nyaa.Search(opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(torrents[0].Name)
	// fmt.Println(torrents[1].Name)
	// fmt.Println(torrents[2].Name)
	// fmt.Println(torrents[3].Name)
	// fmt.Println(torrents[4].Name)

}
