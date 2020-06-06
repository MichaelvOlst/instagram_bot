package bot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

var instagramURL = "https://www.instagram.com"

// Bot holds all config to connect and make request to Instagram
type Bot struct {
	username string
	password string
	profile  string
}

// New returns a new instance of the Bot
func New(username, password, profile string) *Bot {
	return &Bot{
		username,
		password,
		profile,
	}
}

// GetPostUrls gets all the urls from instagram page
func (b *Bot) GetPostUrls() ([]string, error) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	actx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	var urlNodes []*cdp.Node
	err := chromedp.Run(ctx, b.login(), b.getUserFeed(&urlNodes))
	if err != nil {
		return nil, err
	}

	if len(urlNodes) == 0 {
		return nil, errors.New(`URL's are empty for profile ` + b.profile)
	}

	var urls []string
	for _, v := range urlNodes {
		url := fmt.Sprintf("%s%s", instagramURL, v.AttributeValue("href"))
		urls = append(urls, url)
	}

	return urls, nil
}

func (b *Bot) login() chromedp.Tasks {

	emailField := `//input[@name="username"]`
	passwordField := `//input[@name="password"]`

	return chromedp.Tasks{
		chromedp.Navigate(instagramURL),
		chromedp.WaitVisible(emailField),
		chromedp.SendKeys(emailField, b.username),
		chromedp.WaitVisible(passwordField),
		chromedp.SendKeys(passwordField, b.password),
		chromedp.Click(`//button[@type="submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second * 5),
	}
}

func (b *Bot) getUserFeed(urls *[]*cdp.Node) chromedp.Tasks {

	profileURL := fmt.Sprintf("%s/%s/", instagramURL, b.profile)

	return chromedp.Tasks{
		chromedp.Navigate(profileURL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),

		chromedp.Sleep(time.Second * 5),
		chromedp.Nodes(`//a[contains(@href, "/p/")]`, urls),
		chromedp.Sleep(time.Second * 3),
	}
}

// chromedp.Sleep(time.Second * 2),
// chromedp.ActionFunc(func(ctx context.Context) error {
// 	_, exp, err := runtime.Evaluate(`window.scrollTo(0, document.querySelector("footer").offsetTop);`).Do(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	if exp != nil {
// 		return exp
// 	}
// 	return nil
// }),
// chromedp.Sleep(time.Second * 3),
// chromedp.ActionFunc(func(ctx context.Context) error {
// 	_, exp, err := runtime.Evaluate(`window.scrollTo(0, document.querySelector("footer").offsetTop);`).Do(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	if exp != nil {
// 		return exp
// 	}
// 	return nil
// }),
