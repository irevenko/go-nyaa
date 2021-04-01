# 🐈📦 go-nyaa

> nyaa.si client library for Go

<p align="center"><img src="https://avatars.githubusercontent.com/u/28658394?s=200&v=4"></p>

<p align="center">Built on top of <a href="https://github.com/mmcdole/gofeed">gofeed</a>
<p align="center">Original idea <a href="https://github.com/ejnshtein/nyaa-api">ejnshtein/nyaa-api</a></p>


# Installation 🔨
```go get github.com/mmcdole/gofeed``` <br>
```go get github.com/irevenko/go-nyaa```

# Contributing 🤝
Contributions, issues and feature requests are welcome! 👍 <br>
Feel free to check [open issues](https://github.com/irevenko/go-nyaa/issues).

# Docs 📋
Go Reference: https://pkg.go.dev/github.com/irevenko/go-nyaa

```go
import ( 
    "fmt"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	opt := nyaa.SearchOptions{
		Query:    "LN",
		Category: "literature",
		SortBy:   "seeders",
        Filter:   "trusted-only",
	}

	torrents, err := nyaa.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(torrents)
}
```

# Notes
- Pagination does not work with RSS
- Ascending sort does not work with RSS

# ToDo
- [x] Sort by comments, size, date, seeders, leechers,  
- [x] Filters: No filters, No Remakes, Trusted only 
- [x] Search Home Page (basically all newest torrents)
- [x] Search Anime
- [x] Search Literature (Mostly Mangas,LNs)
- [x] Search Audio
- [x] Search Live Action
- [x] Search Pictures
- [x] Search Software (apps, games)
- [ ] Add sukebei.nyaa.si support
- [ ] Scrap full description and comments from nyaa.si/view/...

# What I Learned 🧠
- RSS Feed, xml

# License 📑 
(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)