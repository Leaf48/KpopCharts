package pkg

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type (
	chart struct {
		title []string
		group []string
	}

	Group struct {
		Group string `json:"group"`
		Title string `json:"title"`
	}

	groupsList []Group
)

var AllGroups groupsList

func Billboard() groupsList {

	res, err := http.Get("https://www.billboard.com/charts/billboard-korea-100/")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	charts := chart{}

	doc.Find(".pmc-paywall li h3#title-of-a-story").Each(func(i int, s *goquery.Selection) {
		t := strings.TrimSpace(s.Text())
		charts.title = append(charts.title, t)
	})

	doc.Find(".o-chart-results-list-row-container ul li.lrv-u-width-100p ul.lrv-a-unstyle-list li.o-chart-results-list__item span.c-label.a-no-trucate.a-font-primary-s").Each(func(i int, s *goquery.Selection) {
		t := strings.TrimSpace(s.Text())
		charts.group = append(charts.group, t)
	})

	for i := 0; i < len(charts.group); i++ {
		AllGroups = append(AllGroups, Group{Group: charts.group[i], Title: charts.title[i]})
	}
	return AllGroups
}
