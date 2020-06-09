package meta

import (
	"fmt"
	"log"
	"os/exec"
)

type Exiv struct {
	Base string
}

func (e Exiv) runCommand(command string, filename string) {
	cmd := exec.Command(e.Base, "-M", command, filename)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalln("Couldn't execute the exiv command "+command, err)
		fmt.Println(string(stdout))
		return
	}

	fmt.Println(string(stdout))
}

func (e Exiv) WriteTitle(filename string, title string) {
	e.runCommand("set Iptc.Application2.Caption "+title, filename)
	e.runCommand("set Iptc.Application2.Headline "+title, filename)
	// e.runCommand("set Iptc.Application2.Title "+title, filename)
	// e.runCommand("set Exif.Image.ImageDescription "+title, filename)
}

func (e Exiv) WriteDescription(filename string, description string) {
	e.runCommand("set Iptc.Application2.ObjectName "+description, filename)
}

func (e Exiv) WriteKeywords(filename string, keywords []string) {
	e.runCommand("del Iptc.Application2.Keywords", filename)
	for _, word := range keywords {
		e.runCommand("add Iptc.Application2.Keywords "+word, filename)
	}
}

func NewExiv() *Exiv {
	return &Exiv{"exiv2"}
}

type Attrs struct {
	Title       *string
	Description *string
	Keywords    []string
}

func WriteAttrs(jpegFileName string, attrs Attrs) {
	exiv := NewExiv()
	if attrs.Title != nil {
		exiv.WriteTitle(jpegFileName, *attrs.Title)
	}
	if attrs.Description != nil {
		exiv.WriteDescription(jpegFileName, *attrs.Description)
	}
	if attrs.Keywords != nil {
		exiv.WriteKeywords(jpegFileName, attrs.Keywords)
	}
}
