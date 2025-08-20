package qoutes

import (
	"math/rand"
	"time"
)
type Quote struct{
	text string
}
var Quotes = []Quote{
	{text: "Stay hungry, stay foolish."},
	{text: "Simplicity is the ultimate sophistication."},
	{text: "Code is like humor. When you have to explain it, it’s bad."},
	{text: "Great things never come from comfort zones."},
	{text: "Do what you can with what you have, where you are."},
}

func GetRandomQuote() Quote {
	rand.Seed(time.Now().UnixNano())
	return Quotes[rand.Intn(len(Quotes))]
}