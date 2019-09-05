package login

import (
	"context"
	"github.com/chromedp/chromedp"
	"lotterySpider/processor"
)

const (
	weiboLogginPage = "https://weibo.com/"
)

type LoginProcessor interface {
	Login(account string, password string) error
}

type githubLoginProcessor struct {
}

func NewGithubLoginProcessor() githubLoginProcessor {
	return githubLoginProcessor{}
}

func (g githubLoginProcessor) Login(ctx context.Context, account string, password string) error {

	if err := chromedp.Run(ctx, chromedp.Navigate("https://github.com/login")); err != nil {
		return err
	}
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`//input[@type="submit"]`)); err != nil {
		return err
	}
	return processor.Submit(ctx,"//form", &processor.InputBoxContent{
		InputBoxIndex: `//input[@name="login"]`,
		InputBoxText:  account,
	}, &processor.InputBoxContent{
		InputBoxIndex: `//input[@name="password"]`,
		InputBoxText:  password,
	})
}

type weiboLoginProcessor struct {
}

func NewWeiboLoginProcessor() weiboLoginProcessor {
	return weiboLoginProcessor{}
}

func (w weiboLoginProcessor) Login(ctx context.Context, account string, password string) error {
	tasks := chromedp.Tasks{
		chromedp.Navigate("https://www.weibo.com/login.php"),
		//chromedp.WaitVisible(`//input[@id="loginname"]`),

	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	return processor.Submit(ctx,"//form", &processor.InputBoxContent{
		InputBoxIndex: `//input[@id="loginname"]`,
		InputBoxText:  account,
	}, &processor.InputBoxContent{
		InputBoxIndex: `//input[@name="password"]`,
		InputBoxText:  password,
	})
}
