package main

import (
	"fmt"

	"github.com/Zereker/scraper/pkg/twitter"
)

func main() {
	scraper := twitter.New()
	if err := scraper.LoginOpenAccount(); err != nil {
		panic(err)
	}

	trends, err := scraper.GetTrends()
	if err != nil {
		panic(err)
	}

	fmt.Println(trends)
}
