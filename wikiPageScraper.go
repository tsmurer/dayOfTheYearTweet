package main

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func scrape(htmlText string) []string {
	tkn := html.NewTokenizer(strings.NewReader(htmlText))
	compiler := regexp.MustCompile(`(AD )?\d{1,4}( BC)?( +)– .+`)
	isLiTag := false
	factsList := []string{}
	factString := ""
	isInMain := false
	for {
		tt := tkn.Next()
		t := tkn.Token()
		switch {
		case tt == html.ErrorToken:
			return nil
		case tt == html.StartTagToken:

			if isInBodyContent(t) {
				isInMain = true
			} else if t.Data == "li" && isInMain {
				isLiTag = true
			}
			if isSomethingThatHappened(t) && isInMain {
				factsList = append(factsList, t.Attr[1].Val)
			} else if len(t.Attr) > 1 && t.Attr[1].Val == "Holidays_and_observances" {
				return factsList
			}

		case tt == html.TextToken:
			if isLiTag {
				factString = factString + t.Data
			}
		case tt == html.EndTagToken:
			if t.Data == "main" {
				isInMain = false
			}
			if t.Data == "li" {
				//fmt.Println("factString: ", factString)
				isLiTag = false
				matchesString := compiler.Match([]byte(factString))
				if matchesString {
					factsList = append(factsList, factString)
					factString = ""
				}
			}
		}
	}
}

func isSomethingThatHappened(t html.Token) bool {
	happenings := "Events Births Deaths"
	return t.Data == "span" && len(t.Attr) > 1 && strings.Contains(happenings, string(t.Attr[1].Val))
}

func isInBodyContent(t html.Token) bool {
	return len(t.Attr) > 0 && t.Attr[0].Val == "bodyContent"
}
