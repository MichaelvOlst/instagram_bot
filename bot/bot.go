package bot

import (
	"context"
	"errors"
	"fmt"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var instagramURL = "https://www.instagram.com"

// Bot holds all config to connect and make request to Instagram
type Bot struct {
	Username  string
	Password  string
	Profile   string
	PostURLS  []*cdp.Node
	PostPages []*PostPage
}

// PostPage holds the data when we are visiting a page
type PostPage struct {
	URL    string
	Images []string
	Video  string
	HTML   string
}

// New returns a new instance of the Bot
func New(username, password, profile string) *Bot {
	var postURLS []*cdp.Node
	var PostPages []*PostPage

	return &Bot{
		username,
		password,
		profile,
		postURLS,
		PostPages,
	}
}

// GetPosts gets all the urls from instagram page
func (b *Bot) GetPosts() ([]string, error) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
	actx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	err := chromedp.Run(ctx, b.login(), b.getProfilePosts())
	if err != nil {
		return nil, err
	}

	if len(b.PostURLS) == 0 {
		return nil, errors.New(`URL's are empty for profile ` + b.Profile)
	}

	var urls []string
	for _, v := range b.PostURLS {
		url := fmt.Sprintf("%s%s", instagramURL, v.AttributeValue("href"))
		urls = append(urls, url)
	}

	// var wg sync.WaitGroup
	i := 0
	for _, url := range urls {
		// wg.Add(1)

		// go func(wg *sync.WaitGroup, url string) {
		// 	defer wg.Done()

		b.goToPostDetail(ctx, url)
		// }(&wg, url)

		if i == 2 {
			break
		}
		i++
	}

	// wg.Wait()

	// var images []string
	// for _, v := range b.postDetailImages {
	// 	image := fmt.Sprintf("%s", v.AttributeValue("srcset"))

	// 	// splitted := strings.Split(image, ",")

	// 	// images = append(images, splitted...)
	// 	images = append(images, image)
	// }

	for _, p := range b.PostPages {
		fmt.Printf("%+v\n", p.Images)
		fmt.Println("")
	}

	for _, p := range b.PostPages {
		fmt.Printf("%+v\n", p.Video)
		fmt.Println("")
	}

	// fmt.Printf(b.html)

	return urls, nil
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
