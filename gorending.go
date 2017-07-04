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

// CrawlTrending is function for printing parsed trending
func CrawlTrending(lang string, count int) error {
	doc, err := goquery.NewDocument("https://github.com/trending/" + lang)
	if err != nil {
		log.Fatal(err)
		return err
	}

	repoList := doc.Find(".d-inline-block > h3 > a").Slice(0, count)

	repoList.Each(func(i int, s *goquery.Selection) {
		repoName := strings.Trim(s.Text(), " \n")
		fmt.Printf("%d - %s\n", i+1, repoName)
	})
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gorending"
	app.Usage = "Show Github trending in Terminal!"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2017 Myungseo Kang"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Myungseo Kang",
			Email: "l3opold7@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, L",
			Value: "",
			Usage: "language that you want to see, default all language",
		},
		cli.IntFlag{
			Name:  "count, C",
			Value: 10,
			Usage: "count that you want to see, defalut 10",
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
