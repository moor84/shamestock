package meta

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(url string) *Attrs {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("h1.m_b_b").First().Text()

	var keywords = []string{}
	doc.Find("a.b_aB_a.b_aB_b.b_aB_c").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Text())
		keywords = append(keywords, s.Text())
	})

	return &Attrs{
		Title:       &title,
		Description: &title,
		Keywords:    keywords,
	}
}
