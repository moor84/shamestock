package meta

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WriteAttrsFromCSV(csvFileName string) {
	fmt.Println("Start")

	csvfile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

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
		title := record[1]
		keywords := record[2]

		if filename == "Filename" {
			continue
		}

		filename = strings.ReplaceAll(filename, ".eps", ".jpg")
		filename = filepath.Dir(csvFileName) + "/" + filename

		fmt.Println(filename)

		WriteAttrs(filename, Attrs{
			Title:       &title,
			Description: &title,
			Keywords:    strings.Split(keywords, ","),
		})
	}

	fmt.Println("Done")
}
