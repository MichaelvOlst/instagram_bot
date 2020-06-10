package bot

import (
	"context"
	"encoding/json"
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

	doc.Find("body script").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		if strings.Contains(t, "graphql") && strings.Contains(t, "GraphVideo") {
			jsonData := t[strings.Index(t, "{") : len(t)-2]

			videoData := &graphqlVideoData{}
			err := json.Unmarshal([]byte(jsonData), videoData)
			if err != nil {
				log.Fatal(err)
			}

			p.VideoData = videoData
		}

		if strings.Contains(t, "graphql") && strings.Contains(t, "GraphImage") {
			jsonData := t[strings.Index(t, "{") : len(t)-2]

			imageData := &graphqlImageData{}
			err := json.Unmarshal([]byte(jsonData), imageData)
			if err != nil {
				log.Fatal(err)
			}

			p.ImageData = imageData
		}
	})

	b.Posts = append(b.Posts, p)

}

func (b *Bot) getPostDetail(url string, p *Post) chromedp.Tasks {

	fmt.Printf("Opening %s\n", url)

	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second * 2),
		chromedp.OuterHTML("body", &p.HTML),
	}

}
