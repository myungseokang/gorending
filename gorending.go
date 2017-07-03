package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gorending"
	app.Usage = "Show Github trending in Terminal!"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "",
			Usage: "language that you want",
		},
	}

	app.Action = func(c *cli.Context) error {
		lang := c.String("lang")

		resp, err := http.Get("https://github.com/trending/" + lang)
		if err != nil {
			fmt.Println("http error is: ", err)
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error is: ", err)
			return err
		}

		bodyStr := string(body)

		fmt.Println(bodyStr)
		fmt.Println("lang:", lang)

		return nil
	}

	app.Run(os.Args)
}
