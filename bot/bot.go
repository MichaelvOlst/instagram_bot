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
	Username string
	Password string
	Profile  string
	Webhook  string
	PostURLS []*cdp.Node
	Posts    []*Post
}

// Post holds the data when we are visiting a page
type Post struct {
	URL         string     `json:"url"`
	HTML        string     `json:"-"`
	Images      *postImage `json:"images"`
	Video       *postVideo `json:"video"`
	Description string     `json:"description"`
}

type postImage struct {
	W640  string `json:"w640"`
	W750  string `json:"w750"`
	W1080 string `json:"w1080"`
}

type postVideo struct {
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
	Views    int    `json:"views"`
	Comments int    `json:"comments"`
}

// InstagramRequest holds the data when an user request the data via the API call
type InstagramRequest struct {
	WebhookURL string `json:"webhook_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Profile    string `json:"profile"`
}

// InstagramResult holds the crawled data
type InstagramResult struct {
	Profile  string  `json:"profile"`
	Username string  `json:"username"`
	Webhook  string  `json:"webhook_url"`
	Posts    []*Post `json:"posts"`
}

// New returns a new instance of the Bot
func New(r *InstagramRequest) *Bot {
	var postURLS []*cdp.Node
	var Posts []*Post

	return &Bot{
		r.Username,
		r.Password,
		r.Profile,
		r.WebhookURL,
		postURLS,
		Posts,
	}
}

// Run will start the crawling process
func (b *Bot) Run() {
	w := &WebhookResponse{}

	result, err := b.GetPosts()
	if err != nil {
		w.Error = err
		b.makeRequest(w)
	}

	w.Response = result
	b.makeRequest(w)
}

// GetPosts gets all the urls from instagram page
func (b *Bot) GetPosts() (*InstagramResult, error) {

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

	i := 0
	for _, url := range urls {
		b.goToPostDetail(ctx, url)
		if i == 2 {
			break
		}
		i++
	}

	return &InstagramResult{
		b.Profile,
		b.Username,
		b.Webhook,
		b.Posts,
	}, nil
}
