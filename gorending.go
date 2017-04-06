package gorending

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Gorending() string {
	resp, err := http.Get("https://github.com/trending")
	if err != nil {
		fmt.Println("http transport error is: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error is: ", err)
	}

	bodyStr := string(body)

	return bodyStr
}
