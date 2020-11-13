package meta

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func loadPage(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return doc
}

func scrapeShutterstockURL(url string) *Attrs {
	doc := loadPage(url)
	if doc == nil {
		return nil
	}

	title := doc.Find("h1[data-automation='ImageDetailsPage_Details']").First().Text()

	var keywords = []string{}
	doc.Find("div[data-automation='ExpandableKeywordsList_container_div'] a").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Text())
		keywords = append(keywords, s.Text())
	})

	return &Attrs{
		Title:       &title,
		Description: &title,
		Keywords:    keywords,
	}
}

func scrapeIStockURL(url string) *Attrs {
	doc := loadPage(url)
	if doc == nil {
		return nil
	}

	title := doc.Find("section.title").First().Text()
	title = strings.Replace(title, "...", ". ", 1)

	var keywords = []string{}
	var kw string
	doc.Find("ul.keywords-links a").Each(func(i int, s *goquery.Selection) {
		kw = s.Text()
		kw = strings.Replace(kw, ", ", "", 1)
		kw = strings.Replace(kw, " Illustrations", "", 1)
		kw = strings.ToLower(kw)
		keywords = append(keywords, kw)
	})

	return &Attrs{
		Title:       &title,
		Description: &title,
		Keywords:    keywords,
	}
}

func Scrape(url string) *Attrs {
	var attrs *Attrs
	if strings.HasPrefix(url, "https://www.shutterstock.com/") {
		attrs = scrapeShutterstockURL(url)
	} else if strings.HasPrefix(url, "https://www.istockphoto.com/") {
		attrs = scrapeIStockURL(url)
	} else {
		log.Fatal("URL is not recognised")
		attrs = nil
	}
	if attrs != nil {
		return attrs
	}
	title := ""
	desc := ""
	var keywords = []string{}
	return &Attrs{
		Title:       &title,
		Description: &desc,
		Keywords:    keywords,
	}
}

func ScrapeAttrsCSV(csvFileName string) {
	fmt.Println("Start")

	csvfile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	outFileName := filepath.Dir(csvFileName) + "/attrs.csv"
	outfile, err := os.Create(outFileName)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}
	writer := csv.NewWriter(outfile)

	reader := csv.NewReader(csvfile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		filename := record[0]
		url := record[1]

		if filename == "Filename" {
			continue
		}

		attrs := Scrape(url)

		keywords := strings.Join(attrs.Keywords, ",")
		newRecord := []string{filename, *attrs.Title, keywords, url}

		fmt.Println(newRecord)

		if err := writer.Write(newRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func CreateCSVtemplate(dir string) {
	outFileName := dir + "/urls.csv"
	outfile, err := os.Create(outFileName)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}
	writer := csv.NewWriter(outfile)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".eps") {
			continue
		}

		newRecord := []string{file.Name(), ""}

		if err := writer.Write(newRecord); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
