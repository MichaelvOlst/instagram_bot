package bot

import (
	"time"

	"github.com/chromedp/chromedp"
)

func (b *Bot) login() chromedp.Tasks {

	emailField := `//input[@name="username"]`
	passwordField := `//input[@name="password"]`

	return chromedp.Tasks{
		chromedp.Navigate(instagramURL),
		chromedp.WaitVisible(emailField),
		chromedp.SendKeys(emailField, b.Username),
		chromedp.WaitVisible(passwordField),
		chromedp.SendKeys(passwordField, b.Password),
		chromedp.Click(`//button[@type="submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second * 5),
	}
}
