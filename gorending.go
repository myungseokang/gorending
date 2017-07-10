// Gorending is cli tool that crawls github trendings in Terminal at real time.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

// GithubURL is Github url
var GithubURL = "https://github.com"

// TrendingURL is Github trending url
var TrendingURL = GithubURL + "/trending/"

// CrawlTrending function requests to Github website and prints parsed trendings.
func CrawlTrending(lang string, count int) error {
	doc, err := goquery.NewDocument(TrendingURL + lang)
	if err != nil {
		log.Fatal(err)
		return err
	}

	repoList := doc.Find(".d-inline-block > h3 > a").Slice(0, count)

	repoList.Each(func(i int, s *goquery.Selection) {
		repoURL, ok := s.Attr("href")
		repoName := strings.Trim(s.Text(), " \n")
		if ok {
			fmt.Printf("%d - %s (%s)\n", i+1, repoName, GithubURL+repoURL)
		}
	})
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gorending"
	app.Usage = "Show Github trending in Terminal!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2017 Myungseo Kang"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Myungseo Kang",
			Email: "l3opold7@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		// lang flag is language that you want to see
		cli.StringFlag{
			Name:  "lang, L",
			Value: "",
			Usage: "language that you want to see (default: all language)",
		},
		// count flag is count that you want to see
		cli.IntFlag{
			Name:  "count, C",
			Value: 10,
			Usage: "count that you want to see",
		},
	}

	app.Action = func(c *cli.Context) error {
		lang := c.String("lang")
		count := c.Int("count")

		err := CrawlTrending(lang, count)
		if err != nil {
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
