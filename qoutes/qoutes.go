package qoutes

import (
	"math/rand"
	"time"
)
var Quotes = []string{
	"Stay hungry, stay foolish.",
	"Simplicity is the ultimate sophistication.",
	"Code is like humor. When you have to explain it, it’s bad.",
	"Great things never come from comfort zones.",
	"Do what you can with what you have, where you are.",
}



func GetRandomQuote() string {
	rand.Seed(time.Now().UnixNano())
	return Quotes[rand.Intn(len(Quotes))]
}