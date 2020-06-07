package bot

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (b *Bot) getProfilePosts() chromedp.Tasks {

	profileURL := fmt.Sprintf("%s/%s/", instagramURL, b.Profile)

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
		chromedp.Nodes(`//a[contains(@href, "/p/")]`, &b.PostURLS),
		chromedp.Sleep(time.Second * 1),
	}
}
