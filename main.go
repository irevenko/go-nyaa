package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://nyaa.si/?page=rss")

	var res []Torrent
	for _, item := range feed.Items {
		res = append(
			res,
			Torrent{
				Name:        item.Title,
				Link:        item.Link,
				Date:        item.Published,
				Description: item.Description,
				GUID:        item.GUID,
				Comments:    item.Extensions["nyaa"]["comments"][0].Value,
				Trusted:     item.Extensions["nyaa"]["trusted"][0].Value,
				Remake:      item.Extensions["nyaa"]["remake"][0].Value,
				Size:        item.Extensions["nyaa"]["size"][0].Value,
				Seeders:     item.Extensions["nyaa"]["seeders"][0].Value,
				Leechers:    item.Extensions["nyaa"]["leechers"][0].Value,
				Downloads:   item.Extensions["nyaa"]["downloads"][0].Value,
				Category:    item.Extensions["nyaa"]["category"][0].Value,
				CategoryID:  item.Extensions["nyaa"]["categoryId"][0].Value,
				InfoHash:    item.Extensions["nyaa"]["infoHash"][0].Value,
			},
		)
	}

	fmt.Println(len(res))
}

type Torrent struct {
	Category    string
	Name        string
	Description string
	Date        string
	Size        string
	Seeders     string
	Leechers    string
	Downloads   string
	Trusted     string
	Remake      string
	Comments    string
	Link        string
	GUID        string
	CategoryID  string
	InfoHash    string
}
