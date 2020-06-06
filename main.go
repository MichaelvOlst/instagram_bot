package main

import (
	"fmt"
	"instagram_bot/bot"
	"log"
)

func main() {

	b := bot.New(username, password, profilePage)
	urls, err := b.GetPostUrls()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", len(urls))

	for _, v := range urls {
		fmt.Printf("%s\n", v)
	}
}
