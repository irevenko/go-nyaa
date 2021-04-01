package nyaa

import "fmt"

const (
	baseURL         = "https://nyaa.si/?page=rss&q=+"
	sortByComments  = "&s=comments&o=desc"
	sortBySeeders   = "&s=seeders&o=desc"
	sortByLeechers  = "&s=leechers&o=desc"
	sortByDownloads = "&s=downloads&o=desc"
	sortBySize      = "&s=size&o=desc"
	sortByDate      = "&s=id&o=desc"

	filterNoFilters   = "&f=0"
	filterNoRemakes   = "&f=1"
	filterTrustedOnly = "&f=2"

	categoryAll        = "&c=0_0"
	categoryAnime      = "&c=1_0"
	categoryAudio      = "&c=2_0"
	categoryLiterature = "&c=3_0"
	categoryLiveAction = "&c=4_0"
	categoryPictures   = "&c=5_0"
	categorySoftware   = "&c=6_0"
)

func buildURL(opts SearchOptions) (string, error) {
	url := baseURL

	if opts.Query != "" {
		url += opts.Query
	}

	if opts.Category != "" {
		switch opts.Category {
		case "all":
			url += categoryAll
		case "anime":
			url += categoryAnime
		case "audio":
			url += categoryAudio
		case "literature":
			url += categoryLiterature
		case "live-action":
			url += categoryLiveAction
		case "pictures":
			url += categoryPictures
		case "software":
			url += categorySoftware
		default:
			err := fmt.Errorf("such category option does not exitst\nsee docs: https://github.com/irevenko/go-nyaa")
			return "", err
		}
	}

	if opts.SortBy != "" {
		switch opts.SortBy {
		case "downloads":
			url += sortByDownloads
		case "comments":
			url += sortByComments
		case "seeders":
			url += sortBySeeders
		case "leechers":
			url += sortByLeechers
		case "size":
			url += sortBySize
		case "date":
			url += sortByDate
		default:
			err := fmt.Errorf("such sort option does not exitst\nsee docs: https://github.com/irevenko/go-nyaa")
			return "", err
		}
	}

	if opts.Filter != "" {
		switch opts.Filter {
		case "no-filters":
			url += filterNoFilters
		case "no-remakes":
			url += filterNoRemakes
		case "trusted-only":
			url += filterTrustedOnly
		default:
			err := fmt.Errorf("such filter option does not exitst\nsee docs: https://github.com/irevenko/go-nyaa")
			return "", err
		}
	}

	return url, nil
}
