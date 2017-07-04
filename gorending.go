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

func CrawlTrending(lang string) error {
	doc, err := goquery.NewDocument("https://github.com/trending/" + lang)
	if err != nil {
		log.Fatal(err)
		return err
	}

	doc.Find(".d-inline-block > h3 > a").Each(func(i int, s *goquery.Selection) {
		repoName := strings.Trim(s.Text(), " \n")
		fmt.Printf("%d - %s\n", i, repoName)
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
			Name:  "lang, l",
			Value: "",
			Usage: "language that you want",
		},
	}

	app.Action = func(c *cli.Context) error {
		lang := c.String("lang")

		err := CrawlTrending(lang)
		if err != nil {
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
