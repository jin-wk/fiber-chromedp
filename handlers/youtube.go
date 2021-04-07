package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber"
	"github.com/jin-wk/fiber-chromedp/database"
	"github.com/jin-wk/fiber-chromedp/models"
)

// Crawl godoc
// @Summary Crawl Youtube Comments
// @Description Crawl Youtube Comments
// @Tags crawl
// @Accept  json
// @Produce  json
// @Param url query string true "url"
// @Success 200 {object} models.ResponseModel{}
// @Failure 404 {object} models.ResponseModel{}
// @Failure 503 {object} models.ResponseModel{}
// @Router /api/crawl/youtube [get]
func Crawl(c *fiber.Ctx) error {
	var err error

	url := c.Query("url")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless", false),
	)
	alloCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(
		alloCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*300)
	defer cancel()

	start := time.Now()
	err = chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*3),
	)
	handleError(err)

	height, err := GetScrollHeight(ctx)
	handleError(err)

	var currentHeight int
	for {
		err = chromedp.Run(
			ctx,
			chromedp.ActionFunc(func(c context.Context) error {
				_, exp, err := runtime.Evaluate(`window.scrollTo(0, document.documentElement.scrollHeight)`).Do(c)
				if err != nil {
					return err
				}
				if exp != nil {
					return exp
				}
				return nil
			}),
			chromedp.Sleep(time.Second*1),
			chromedp.Evaluate(`document.documentElement.scrollHeight`, &currentHeight),
		)
		handleError(err)

		if height == currentHeight {
			break
		}
		height = currentHeight
	}

	var html string
	err = chromedp.Run(
		ctx,
		chromedp.OuterHTML(`ytd-comments#comments`, &html),
	)
	handleError(err)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	handleError(err)

	var ids, comments []string
	doc.Find("div#header-author > h3 > a#author-text > span").Each(func(i int, s *goquery.Selection) {
		ids = append(ids, CleanString(s.Text()))
	})
	doc.Find("div#content > yt-formatted-string#content-text").Each(func(i int, s *goquery.Selection) {
		comments = append(comments, CleanString(s.Text()))
	})

	var commentModel []models.Comment
	for k, v := range ids {
		commentModel = append(
			commentModel,
			models.Comment{
				UserID:  v,
				Comment: comments[k],
			},
		)
	}

	err = database.DB.Create(&commentModel).Error
	handleError(err)

	fmt.Printf("\nTook: %.3f secs\n", time.Since(start).Seconds())

	return c.JSON(models.Response(1000, "success", nil))
}

func GetScrollHeight(ctx context.Context) (int, error) {
	var height int
	err := chromedp.Run(ctx, chromedp.Evaluate(`document.documentElement.scrollHeight`, &height))
	return height, err
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
