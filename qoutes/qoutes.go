package qoutes

import (
	"math/rand"
	"time"
)

type Quote struct {
	Text string `json:"text"`
}

var Quotes = []Quote{
	{Text: "Stay hungry, stay foolish."},
	{Text: "Simplicity is the ultimate sophistication."},
	{Text: "Code is like humor. When you have to explain it, it’s bad."},
	{Text: "Great things never come from comfort zones."},
	{Text: "Do what you can with what you have, where you are."},
}

func GetRandomQuote() Quote {
	rand.Seed(time.Now().UnixNano())
	return Quotes[rand.Intn(len(Quotes))]
}
