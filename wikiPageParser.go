package main

import (
	"strings"

	"golang.org/x/net/html"
)

func parse(htmlText string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(htmlText))
	var vals []string
	var typeOfThingThatHappened string
	isLiTag := false
	eventList := []Event{}
	birthList := []Person{}
	deathList := []Person{}
	holidayList := []Holiday{}
	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()
			if isSomethingThatHappened(t) {
				typeOfThingThatHappened = t.Attr[1].Val
			}

			if t.Data == "li" {
				isLiTag = true
			}

		case tt == html.TextToken:

			t := tkn.Token()
			if isLiTag && typeOfThingThatHappened != "" {
				text := strings.Split(t.Data, " - ")
				year, descr := text[0], text[1]
				switch typeOfThingThatHappened {
				case "Events":
					thing := ThingThatHappened{
						year,
						descr,
					}
					eventList = append(eventList, Event{thing})

				case "Births":
					descr := strings.SplitN(descr, ",", 2)
					name, who := descr[0], descr[1]
					thing := ThingThatHappened{
						year,
						who,
					}
					birthList = append(birthList, Person{thing, name})

				case "Deaths":
					descr := strings.SplitN(descr, ",", 2)
					name, who := descr[0], descr[1]
					thing := ThingThatHappened{
						year,
						who,
					}
					deathList = append(deathList, Person{thing, name})
				case "Holidays_and_observances":
					holiday := Holiday{year}
					holidayList = append(holidayList, holiday)
				}
			}
			isLiTag = false

		}
	}
}

func isSomethingThatHappened(t html.Token) bool {
	happenings := "Events Births Deaths Holidays_and_observances"
	return t.Data == "span" && len(t.Attr) > 1 && strings.Contains(happenings, string(t.Attr[1].Val))
}
