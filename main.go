package main

import (
	"instagram_bot/server"
)

func main() {

	s := server.New()

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
