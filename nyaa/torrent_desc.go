package nyaa

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

func TorrentDescription(id int, provider string) (string, error) {
	var description string

	c := colly.NewCollector()

	c.OnHTML("#torrent-description", func(e *colly.HTMLElement) {
		description = e.Text
	})

	var e error
	c.OnError(func(r *colly.Response, err error) {
		e = err
	})
	if e != nil {
		return "", e
	}

	if provider == "nyaa" {
		c.Visit(nyaaView + strconv.Itoa(id))
	} else if provider == "sukebei" {
		c.Visit(sukebeiView + strconv.Itoa(id))
	} else {
		err := fmt.Errorf("provider param could be nyaa or sukebei\nsee docs: https://github.com/irevenko/go-nyaa#provider")
		return "", err
	}

	return description, nil
}
