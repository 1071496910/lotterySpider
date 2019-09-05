package processor

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
)

type InputBoxContent struct {
	InputBoxIndex string
	InputBoxText  string
}

func Submit(ctx context.Context, formIndex string, contents ...*InputBoxContent) error {

	if err := chromedp.Run(ctx, chromedp.WaitVisible(formIndex)); err != nil {
		return err
	}
	if len(contents) == 0 {
		return chromedp.Run(ctx, chromedp.Submit(formIndex))
	}
	for _, content := range contents {
		tasks := chromedp.Tasks{
			chromedp.WaitVisible(content.InputBoxIndex),
			chromedp.SendKeys(content.InputBoxIndex, content.InputBoxText),
		}
		if err := chromedp.Run(ctx, tasks); err != nil {
			return err
		}
	}
	return chromedp.Run(ctx, chromedp.Submit(formIndex))
}

func DebugCapScreen(ctx context.Context, file string) {
	var res []byte
	if err := chromedp.Run(ctx, chromedp.CaptureScreenshot(&res)); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(file, res, 0644); err != nil {
		panic(err)
	}
}
