package nyaa

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	t "github.com/irevenko/go-nyaa/types"
)

func TorrentComments(id int, provider string) ([]t.Comment, error) {
	var comments []t.Comment

	c := colly.NewCollector()

	c.OnHTML(".comment-panel", func(e *colly.HTMLElement) {
		comments = append(comments, t.Comment{
			User: e.ChildText(".text-default"),
			Date: e.ChildText(".comment-details"),
			Text: e.ChildText(".comment-body"),
		})
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

	return comments, nil
}
