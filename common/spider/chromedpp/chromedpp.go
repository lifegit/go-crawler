/**
* @Author: TheLife
* @Date: 2021/5/10 下午3:56
 */
package chromedpp

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type ChromeDpp struct {
	c       *context.Context
	cancels *[]context.CancelFunc
}

// create chrome instance
func NewChromeDp(timeOutSecond time.Duration, headless bool) (*context.Context, context.CancelFunc) {
	// 1.chrome conf
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		//chromedp.ProxyServer("http://username:password@proxyserver.com:31280"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	// 2.NewExecAllocator
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	//defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	//defer cancel()

	// 3. create a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	return &ctx, cancel
}
