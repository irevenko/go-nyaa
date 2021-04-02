package nyaa

import (
	"github.com/gocolly/colly"
	t "github.com/irevenko/go-nyaa/types"
)

func TorrentComments(viewURL string) ([]t.Comment, error) {
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

	c.Visit(viewURL)

	return comments, nil
}
