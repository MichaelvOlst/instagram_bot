package main

import (
	"fmt"
	"instagram_bot/config"
	"instagram_bot/database"
	"instagram_bot/server"
	"log"
	"path/filepath"
)

func main() {

	envFile, _ := filepath.Abs(".env")
	cfg := config.Parse(envFile)

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("%+v\n", cfg)

	s := server.New(cfg, db)

	s.Start()

	// done := make(chan bool)

	// b := bot.New(username, password, profilePage)

	// go func(done chan bool) {

	// 	urls, err := b.GetPosts()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, result := range urls {
	// 		fmt.Println(result)
	// 		// fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	// 	}

	// 	done <- true
	// }(done)

	// <-done

}
