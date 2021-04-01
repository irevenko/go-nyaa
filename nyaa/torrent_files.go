package nyaa

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

func TorrentFiles(id int, provider string) ([]string, error) {
	var folders []string
	var files []string

	c := colly.NewCollector()

	c.OnHTML(".folder", func(e *colly.HTMLElement) {
		folders = append(folders, e.Text)
	})

	c.OnHTML(".torrent-file-list", func(e *colly.HTMLElement) {
		files = append(files, e.ChildText("li"))
	})

	var e error
	c.OnError(func(r *colly.Response, err error) {
		e = err
	})
	if e != nil {
		return nil, e
	}

	if provider == "nyaa" {
		c.Visit(nyaaView + strconv.Itoa(id))
	} else if provider == "sukebei" {
		c.Visit(sukebeiView + strconv.Itoa(id))
	} else {
		err := fmt.Errorf("provider param could be nyaa or sukebei\nsee docs: https://github.com/irevenko/go-nyaa#provider")
		return nil, err
	}

	if len(folders) == 0 {
		return files, nil
	}

	return folders, nil
}
