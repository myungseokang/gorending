package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
	"os"
)

const Version = "0.0.1"
const Usage = `
NAME:
    gorending - Show github trending in Terminal!

USAGE:
    gorending [language]

VERSION:
    0.0.1

COMMANDS:
    help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS
    --version Shows version information
    --usage Shows usages
`

func main() {
	lang := flag.String("lang", "", "Language that you want")
	version := flag.Bool("version", false, "Shows version information")
	usage := flag.Bool("usage", false, "Shows usages")

	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	} else if *usage {
		fmt.Println(Usage)
		os.Exit(0)
	}

	resp, err := http.Get("https://github.com/trending/" + *lang)
	if err != nil {
		fmt.Println("http transport error is: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error is: ", err)
	}

	bodyStr := string(body)

	fmt.Println(bodyStr)
	fmt.Println("lang:", *lang)
}
