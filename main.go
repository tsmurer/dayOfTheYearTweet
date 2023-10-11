package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func getDay() string {
	t := time.Now()
	month := t.Month()
	day := t.Day()
	return month.String() + "_" + strconv.Itoa(day)
}

func buildWikiUrl() string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", getDay())
}

func makeWikiRequest() (*http.Response, error) {
	wikiUrl := buildWikiUrl()
	res, err := http.Get(wikiUrl)
	if err != nil {
		return nil, fmt.Errorf("makeWikiRequest: %v", err)
	}
	return res, err
}

func getResponseBody(response *http.Response) (string, error) {
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("getResponseBody: %v", err)
	}
	return string(resBody), err
}

func main() {
	response, err := makeWikiRequest()
	if err != nil {
		fmt.Println(err)
		return
	}

	responseBody, errBody := getResponseBody(response)
	if err != nil {
		fmt.Println(errBody)
		return
	}

	parse(responseBody)

}
