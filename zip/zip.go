package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func CreateZipFiles(dir string) {
	fmt.Println("Start")

	outDir := dir + "/zip"
	err := os.Mkdir(outDir, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".eps") {
			continue
		}
		// fmt.Println(file.Name())
		// fullPath := dir + "/" + file.Name()

		previewName := strings.ReplaceAll(file.Name(), ".eps", ".jpg")
		zipName := strings.ReplaceAll(file.Name(), ".eps", ".zip")

		files := []string{dir + "/" + file.Name(), dir + "/" + previewName}

		if err := ZipFiles(outDir+"/"+zipName, files); err != nil {
			panic(err)
		}
		fmt.Println("Zipped File:", zipName)
	}

	fmt.Println("Done")
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
// https://golangcode.com/create-zip-files-in-go/
func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = path.Base(filename)

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
