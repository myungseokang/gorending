// Gorending is cli tool that crawls github trendings in Terminal at real time.
package gorending

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/manifoldco/promptui"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

// GithubURL is Github url
var GithubURL = "https://github.com"

// TrendingURL is Github trending url
var TrendingURL = GithubURL + "/trending/"

// Repository Structure
type Repository struct {
	URL    string
	Name   string
	Author string
}

// CrawlTrending function requests to Github website and prints parsed trendings.
func CrawlTrending(lang string, count int) []Repository {
	doc, err := goquery.NewDocument(TrendingURL)
	if err != nil {
		log.Fatal(err)
	}

	rawRepoList := doc.Find(".d-inline-block > h3 > a").Slice(0, 10)

	repoList := []Repository{}

	rawRepoList.Each(func(i int, s *goquery.Selection) {
		repoURL, ok := s.Attr("href")
		if ok {
			trendingName := strings.Trim(s.Text(), " \n")
			trendingNameList := strings.Split(trendingName, " / ")
			authorName := trendingNameList[0]
			repoName := trendingNameList[1]

			repo := Repository{
				URL:    GithubURL + repoURL,
				Name:   repoName,
				Author: authorName,
			}
			repoList = append(repoList, repo)
		}
	})

	return repoList
}

func ShowPrompt(repoList []Repository) error {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Name | cyan }} ({{ .URL | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .URL | red }})",
		Selected: "! {{ .Name | red | cyan }}",
		Details: `
--------- Repository ----------
{{ "URL:" | faint }}	 {{ .URL }}
{{ "Name:" | faint }}	 {{ .Name }}
{{ "Author: " | faint }}	 {{ .Author }}`,
	}

	searcher := func(input string, index int) bool {
		repo := repoList[index]
		name := strings.Replace(strings.ToLower(repo.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Repo",
		Items:     repoList,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return err
	}

	open.Run(repoList[i].URL)
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

		repoList := CrawlTrending(lang, count)
		err := ShowPrompt(repoList)

		if err != nil {
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
