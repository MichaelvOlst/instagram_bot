package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func (b *Bot) goToPostDetail(ctx context.Context, url string) {

	var p = &Post{}
	p.URL = url

	err := chromedp.Run(ctx, b.getPostDetail(url, p))
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(p.HTML))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(`article img`).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// band := s.Find("a").Text()
		// title := s.Find("i").Text()
		// fmt.Printf("Review %d: %s - %s\n", i, band, title)
		image, found := s.Attr("srcset")
		if found {
			p.Images = append(p.Images, image)
		}

	})

	doc.Find(`article video`).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// band := s.Find("a").Text()
		// title := s.Find("i").Text()
		// fmt.Printf("Review %d: %s - %s\n", i, band, title)
		image, found := s.Attr("poster")
		if found {
			p.Images = append(p.Images, image)
		}

		videoURL, found := s.Attr("src")
		if found {
			p.Video = videoURL
		}

	})

	b.Posts = append(b.Posts, p)

}

func (b *Bot) getPostDetail(url string, p *Post) chromedp.Tasks {

	fmt.Printf("Opening %s\n", url)

	return chromedp.Tasks{
		chromedp.Navigate(url),
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if exp != nil {
		// 		return exp
		// 	}
		// 	return nil
		// }),

		chromedp.Sleep(time.Second * 2),
		chromedp.OuterHTML("#react-root", &p.HTML),
		// chromedp.Nodes(`img[srcset]`, &b.postDetailImages),
		// chromedp.Nodes(`video[poster]`, &b.postDetailImages),
		// chromedp.Text(`(//*`, &b.html),

		// chromedp.Sleep(time.Second * 3),
	}

}
