package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func getResponse(url string) string {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 45*time.Second)
	defer cancel()
	// run task list
	var res string
	var err error

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(`#availability`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func main() {
	var resText string
	var unavailable bool
	var amzn []string
	amzn = append(amzn, "https://www.amazon.de/-/en/dp/B08H93ZRK9/", "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/")

	for i := 0; i < len(amzn); i++ {
		log.Println("Checking: ", amzn[i])
		resText = getResponse(amzn[i])
		log.Println(strings.TrimSpace(resText))
		unavailable = strings.Contains(resText, "unavailable")
		if unavailable == false {
			os.Exit(100 + i) // fail is ok, github will notify us
		}
	}
}
