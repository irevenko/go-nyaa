# 🐈📦 go-nyaa

[![Go Reference](https://pkg.go.dev/badge/github.com/irevenko/go-nyaa.svg)](https://pkg.go.dev/github.com/irevenko/go-nyaa)
[![Go Report Card](https://goreportcard.com/badge/github.com/irevenko/go-nyaa)](https://goreportcard.com/report/github.com/irevenko/go-nyaa)

> nyaa.si client library for Go

<a href="https://github.com/nyaadevs">
<img src="https://avatars.githubusercontent.com/u/28658394?s=200&v=4">
</a>

Built on top of:<br>
<a href="https://github.com/mmcdole/gofeed">gofeed</a> - search using RSS <br>
<a href="https://github.com/gocolly/colly">colly</a> - scrap torrent details page <br> <br>
Original idea: <br> <a href="https://github.com/ejnshtein/nyaa-api">ejnshtein/nyaa-api</a>


# Installation 🔨
```go get github.com/mmcdole/gofeed``` <br>
```go get github.com/gocolly/colly``` <br>
```go get -u github.com/irevenko/go-nyaa```

# Contributing 🤝
Contributions, issues and feature requests are welcome! 👍 <br>
Feel free to check [open issues](https://github.com/irevenko/go-nyaa/issues).

# Docs 📒
<a href="https://pkg.go.dev/github.com/irevenko/go-nyaa">Go reference</a>, 
<a href="https://github.com/irevenko/go-nyaa/tree/main/_examples">Examples</a>
* [Search Example](#search-example "Goto #search-example-")
* [Search Options](#search-options "Goto #search-options-")
    * [Provider](#provider "Goto #provider-")
    * [Query](#query "Goto #query-")
    * [Category](#category "Goto #category-")
        * [nyaa](#nyaa "Goto #nyaa-")
        * [sukebei](#sukebei "Goto #sukebei-")
    * [SortBy](#sortby "Goto #sortby-")
    * [Filter](#filter "Goto #filter-")
* [TorrentComments](#torrentcomments "Goto #torrentcomments-")
* [TorrentDescription](#torrentdescription "Goto #torrentdescription-")
* [TorrentFiles](#torrentfiles "Goto #torrentfiles-")


## Search Example
```Search``` returns <a href="https://github.com/irevenko/go-nyaa/blob/27885a8e6b01043672b9c8866fc7ff80c837f060/types/types.go#L3">```[]Torrent```</a>

```go
import ( 
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
    opt := nyaa.SearchOptions{
        Provider: "nyaa", // Provider is the only required option
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

## Search Options
```go
type SearchOptions struct {
	Provider string
	Query    string
	Category string
	SortBy   string
	Filter   string
}
```

## Provider
- ```nyaa``` - nyaa.si
- ```sukebei``` - sukebei.nyaa.si

## Query
- your desired search string

## Category
### nyaa
* ```all``` - All Categories
* ```anime``` - All Anime
    * ```anime-amv``` 
    * ```anime-eng```
    * ```anime-non-eng```
    * ```anime-raw```
* ```audio``` - All Audio
    * ```audio-lossless``` 
    * ```audio-lossy```
* ```literature``` - All Literature
    * ```literature-eng``` 
    * ```literature-non-eng```
    * ```literature-raw``` 
* ```live-action``` - All Live Action
    * ```live-action-idol-prom``` 
    * ```live-action-eng```
    * ```live-action-non-eng```
    * ```live-action-raw```
* ```pictures``` - All Pictures
    * ```pictures-graphics``` 
    * ```pictures-photos``` 
* ```software``` - All Software
    * ```software-apps``` 
    * ```software-games```

### sukebei
* ```all``` - All Categories
* ```art``` - All Art
    * ```art-anime``` 
    * ```art-doujinshi```
    * ```art-manga```
    * ```art-games```
    * ```art-pictures```
* ```real-life``` - All Real Life
    * ```real-life-photos``` 
    * ```real-life-videos```

## SortBy
- ```comments```
- ```downloads``` 
- ```date``` 
- ```seeders``` 
- ```leechers```
- ```size``` 

## Filter
- ```no-filter```
- ```no-remakes``` 
- ```trusted-only``` 

## TorrentComments
```TorrentComments``` returns <a href="https://github.com/irevenko/go-nyaa/blob/03bee8828c28766b317229ce57ebd0aa9701c37b/types/types.go#L21">```[]Comment```</a>

```go
import ( 
	"fmt"
	"log"
	
	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	comms, err := nyaa.TorrentComments("https://nyaa.si/view/1366002") // nyaa.si or sukebei.nyaa.si
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
```

## TorrentDescription
```TorrentDescription``` returns ```string```

```go
import ( 
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	desc, err := nyaa.TorrentDescription("https://nyaa.si/view/1366002") // nyaa.si or sukebei.nyaa.si
	if err != nil {
		log.Fatal(err)
	}

   	fmt.Println(desc)
}
```

## TorrentFiles
```TorrentFiles``` returns ```[]string```

```go
import ( 
	"fmt"
	"log"

	"github.com/irevenko/go-nyaa/nyaa"
)

func main() {
	files, err := nyaa.TorrentFiles("https://nyaa.si/view/1366002") // nyaa.si or sukebei.nyaa.si
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		fmt.Println(v)
	}
}
```

# Notes
- Pagination does not work with RSS
- Ascending sort does not work with RSS

# Quick Start 🚀
```git clone https://github.com/irevenko/go-nyaa.git``` <br>
```cd go-nyaa``` <br>
```go get -d ./...``` <br>
```go run _examples/nyaa_search.go``` <br>

# What I Learned 🧠
- RSS Feed, xml
- Parsing html in Go

# License 📑 
(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
