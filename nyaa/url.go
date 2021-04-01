package nyaa

import "fmt"

const (
	baseURL         = "https://nyaa.si/?page=rss&q=+"
	sortByComments  = "&s=comments&o=desc"
	sortBySeeders   = "&s=seeders&o=desc"
	sortByLeechers  = "&s=leechers&o=asc"
	sortByDownloads = "&s=downloads&o=desc"
	sortBySizeDesc  = "&s=size&o=desc"
	sortBySizeAsc   = "&s=size&o=asc"
	sortByDateDesc  = "&s=id&o=desc"
	sortByDateAsc   = "&s=id&o=asc"

	filterNoRemakes   = "&f=1"
	filterTrustedOnly = "&f=2"
)

func buildURL(opts SearchOptions) (string, error) {
	url := baseURL

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
		case "size-asc":
			url += sortBySizeAsc
		case "size-desc":
			url += sortBySizeDesc
		case "date-asc":
			url += sortByDateAsc
		case "date-desc":
			url += sortByDateDesc
		default:
			err := fmt.Errorf("such sort option does not exitst\nsee docs: https://github.com/irevenko/go-nyaa")
			return "", err
		}
	}

	if opts.Filter != "" {
		switch opts.Filter {
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
