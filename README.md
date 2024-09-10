Go OpenGraph
===

Parses given html data into Facebook OpenGraph structure.

To download and install this package run:

`go get github.com/fragoe/go-opengraph/opengraph`

*NOTE: if you need to grab as much info from a page as possible consider using [dyatlov/go-htmlinfo](https://github.com/dyatlov/go-htmlinfo)*

The package supports the whole set of OpenGraph properties from [The Open Graph protocol](https://ogp.me/).

## Command line tool

You can also use `opengraph` from CLI.
You can download latest version of `opengraph` for your OS from [Releases](https://github.com/fragoe/go-opengraph/releases).

You can query website endpoints using the tool directly or use it with other tools for your own workflows.

Example usages:

```bash
# download and parse html page
./opengraph https://www.youtube.com/watch\?v\=yhoI42bdwU4
```

```bash
# parse piped html
curl https://www.youtube.com/watch\?v\=yhoI42bdwU4 | ./opengraph
```

```bash
# get video image
./opengraph https://www.youtube.com/watch\?v\=yhoI42bdwU4 | jq '.images[0].url'
```

## Package Methods

 * `NewOpenGraph()` - create a new OpenGraph instance
 * `ProcessHTML(buffer io.Reader) error` - process given html into underlying data structure
 * `ProcessMeta(metaAttrs map[string]string)` - add data to the structure based on meta attributes
 * `ToJSON() (string, error)` - return JSON representation of data or error
 * `String() string` - return JSON representation of structure

Source docs: http://godoc.org/github.com/fragoe/go-opengraph/opengraph

If you just need to parse an OpenGraph data from HTML then method `ProcessHTML` is your needed one.

Example:

```go
package main

import (
  "fmt"
  "strings"

  "github.com/fragoe/go-opengraph/opengraph"
)

func main() {
  html := `<html><head><meta property="og:type" content="article" />
  <meta property="og:title" content="WordPress 4.3 &quot;Billie&quot;" />
  <meta property="og:url" content="https://wordpress.org/news/2015/08/billie/" /></head><body></body></html>`

  og := opengraph.NewOpenGraph()
  err := og.ProcessHTML(strings.NewReader(html))

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("Type: %s\n", og.Type)
  fmt.Printf("Title: %s\n", og.Title)
  fmt.Printf("URL: %s\n", og.URL)
  fmt.Printf("String/JSON Representation: %s\n", og)
}
```

If you have your own parsing engine and just need an intelligent OpenGraph parsing, then `ProcessMeta` is the method you need.
While using this method you don't need to reparse your parsed html again, just feed it with meta atributes as they appear and OpenGraph will be built based on the data.

Example:

```go
package main

import (
	"fmt"
	"strings"

	"github.com/fragoe/go-opengraph/opengraph"
	"golang.org/x/net/html"
)

func main() {
	h := `<html><head><meta property="og:type" content="article" />
  <meta property="og:title" content="WordPress 4.3 &quot;Billie&quot;" />
  <meta property="og:url" content="https://wordpress.org/news/2015/08/billie/" /></head><body></body></html>`

	og := opengraph.NewOpenGraph()

	doc, err := html.Parse(strings.NewReader(h))
	if err != nil {
		fmt.Println(err)
		return
	}

	var parseHead func(*html.Node)
	parseHead = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "meta" {
				m := make(map[string]string)
				for _, a := range c.Attr {
					m[a.Key] = a.Val
				}

				og.ProcessMeta(m)
			}
		}
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				if c.Data == "head" {
					parseHead(c)
					continue
				} else if c.Data == "body" { // OpenGraph is only in head, so we don't need body
					break
				}
			}
			f(c)
		}
	}
	f(doc)

	fmt.Printf("Type: %s\n", og.Type)
	fmt.Printf("Title: %s\n", og.Title)
	fmt.Printf("URL: %s\n", og.URL)
	fmt.Printf("String/JSON Representation: %s\n", og)
}
```
