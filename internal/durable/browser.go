package durable

import (
	"context"
	"github.com/chromedp/chromedp"
)

func Browser(url string, quality int) ([]byte, error) {
	var buf []byte

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	if err := chromedp.Run(ctx, fullScreenshot(url, quality, &buf)); err != nil {
		return nil, err
	}

	return buf, nil
}

func CompressedImage(url string, quality int) ([]byte, error) {
	image, err := Browser(url, quality)

	if err != nil {
		return nil, err
	}

	compressed, err := Compress(image)

	if err != nil {
		return nil, err
	}

	return compressed, nil
}

func fullScreenshot(urlStr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.FullScreenshot(res, quality),
	}
}
