package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type Influencers struct {
	Rank             string
	ContributorName  string
	ContributorTitle string
	Category         []string
	Subscribers      string
	Audience         string
	Authentic        string
	Engagement       string
}

func Parse(url string) (res []Influencers, err error) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {

			r.HTMLDoc.Find("div.row").Each(func(_ int, selection *goquery.Selection) {
				var infl Influencers
				// Извлечение данных
				infl.Rank = selection.Find("div.row-cell.rank span").First().Text()
				infl.ContributorName = selection.Find("div.contributor__name-content").Text()
				infl.ContributorTitle = selection.Find("div.contributor__title").Text()
				selection.Find("div.tag__content").Each(func(i int, selectionCategory *goquery.Selection) {
					infl.Category = append(infl.Category, selectionCategory.Text())
				})
				infl.Subscribers = selection.Find("div.row-cell.subscribers").Text()
				infl.Audience = selection.Find("div.row-cell.audience").Text()
				infl.Authentic = selection.Find("div.row-cell.authentic").Text()
				infl.Engagement = selection.Find("div.row-cell.engagement").Text()
				res = append(res, infl)
			})
		},
	}).Start()
	return
}
