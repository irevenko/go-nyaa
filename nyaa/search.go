package nyaa

import (
	t "github.com/irevenko/go-nyaa/types"
	"github.com/mmcdole/gofeed"
)

type SearchOptions struct {
	Query    string
	Category string
	SortBy   string
	Filter   string
}

var (
	fp = gofeed.NewParser()
)

func Search(opts SearchOptions) ([]t.Torrent, error) {
	url, err := buildURL(opts)
	if err != nil {
		return nil, err
	}

	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	res := convertRSS(feed)

	return res, nil
}
