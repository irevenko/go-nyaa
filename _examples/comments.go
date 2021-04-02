package main

import (
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	comms, err := nyaa.TorrentComments("https://sukebei.nyaa.si/view/3253887")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range comms {
		fmt.Println("user: " + v.User)
		fmt.Println("at: " + v.Date)
		fmt.Println("text: " + v.Text)
		fmt.Println()
	}
}
