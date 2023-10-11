package main

type ThingThatHappened struct {
	Year        string
	Description string
}
type Event struct {
	ThingThatHappened
}

type Person struct {
	ThingThatHappened
	Name string
}

type Holiday struct {
	Name string
}
