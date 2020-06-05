package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const instagramURL = "https://www.instagram.com"
const username = "michaeltje@live.nl"
const password = "Dkkdkvdt-378834"
const profilePage = "cmspecialist"

func main() {
	// create context

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	actx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	// run task list
	// var res string
	err := chromedp.Run(ctx, login(), getUserFeed())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")

	// log.Printf("got: `%s`", strings.TrimSpace(res))
}

func login() chromedp.Tasks {

	emailField := `//input[@name="username"]`
	passwordField := `//input[@name="password"]`

	return chromedp.Tasks{
		chromedp.Navigate(instagramURL),
		chromedp.WaitVisible(emailField),
		chromedp.SendKeys(emailField, username),
		chromedp.WaitVisible(passwordField),
		chromedp.SendKeys(passwordField, password),
		chromedp.Click(`//button[@type="submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second * 10),
	}
}

func getUserFeed() chromedp.Tasks {

	profileURL := fmt.Sprintf("%s/%s/", instagramURL, profilePage)

	return chromedp.Tasks{
		chromedp.Navigate(profileURL),
		chromedp.Sleep(time.Second * 10),
	}
}
