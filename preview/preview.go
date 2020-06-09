package preview

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Imagemagick struct {
	Base string
}

func (e Imagemagick) runCommand(args []string) {
	cmd := exec.Command(e.Base, args...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalln("Couldn't execute the Imagemagick command", err)
		fmt.Println(string(stdout))
		return
	}

	fmt.Print(string(stdout))
}

func (im Imagemagick) CreatePreview(filename string, outFilename string, density int, quality int) {
	// convert -density 150 -units PixelsPerInch -colorspace sRGB "2154.eps" -quality 100  "2154.jpg"
	var args = []string{
		"-density",
		strconv.Itoa(density),
		"-units",
		"PixelsPerInch",
		"-colorspace",
		"sRGB",
		filename,
		"-quality",
		strconv.Itoa(quality),
		outFilename,
	}
	im.runCommand(args)
}

func NewImagemagick() *Imagemagick {
	return &Imagemagick{"convert"}
}

func CreateNormalPreview(filename string) {
	var previewFilename = strings.ReplaceAll(filename, ".eps", ".jpg")
	NewImagemagick().CreatePreview(filename, previewFilename, 150, 100)
}

func CreateBigPreview(filename string) {
	var previewFilename = strings.ReplaceAll(filename, ".eps", ".jpg")
	NewImagemagick().CreatePreview(filename, previewFilename, 350, 100)
}

func CreatePreviews(dir string) {
	fmt.Println("Start")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".eps") {
			continue
		}

		CreateNormalPreview(dir + "/" + file.Name())
		fmt.Println(dir + "/" + file.Name())
	}

	fmt.Println("Done")
}
