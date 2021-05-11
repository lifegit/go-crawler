/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:21
 */
package tasks

import (
	"github.com/chromedp/chromedp"
	"go-crawler/common/mediem"
	"go-crawler/common/mediem/midMiddleware"
	"go-crawler/common/spider/chromedpp"
)

func search() {
	fetch := func(c *mediem.Context) {
		// 1. create chrome instance
		ctx, cancel := chromedpp.NewChromeDp(10, true)
		defer cancel()

		// 2. search baidu
		var example string
		err := chromedp.Run(*ctx,
			chromedp.Navigate(`https://pkg.go.dev/time`),
			// wait for footer element is visible (ie, page is loaded)
			chromedp.WaitVisible(`body > footer`),
			// find and click "Example" link
			chromedp.Click(`#example-After`, chromedp.NodeVisible),
			// retrieve the text of the textarea
			chromedp.Value(`#example-After textarea`, &example),
		)

		c.Result.Err = err
	}

	var me mediem.Context
	me.Use(midMiddleware.Recovery(), fetch, loggerMid).Run()
}
