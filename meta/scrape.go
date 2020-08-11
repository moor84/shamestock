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

	title := doc.Find("h1[data-automation='ImageDetailsPage_Details']").First().Text()

	var keywords = []string{}
	doc.Find("div[data-automation='ExpandableKeywordsList_container_div'] a.o_button_theme_button").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Text())
		keywords = append(keywords, s.Text())
	})

	return &Attrs{
		Title:       &title,
		Description: &title,
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
