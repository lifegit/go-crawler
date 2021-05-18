/**
* @Author: TheLife
* @Date: 2021/5/10 上午10:21
 */
package tasks

import (
	"github.com/chromedp/chromedp"
	"go-crawler/common/koa"
	"go-crawler/common/spider"
	"go-crawler/common/utils"
	"go-crawler/services/pkgGo/constant"
)

var searchTask *koa.Context

func init() {
	searchTask = utils.NewAweMediem(constant.ServiceName, "search", search)
}

func search(k *koa.Context) {
	// 1. create chrome instance
	ctx, cancel := spider.NewAweChromeDp(10, true)
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

	k.Result.Data = example
	k.Result.Err = err
}
