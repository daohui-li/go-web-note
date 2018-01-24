package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var initilized = false

func persistInit() {
	if initilized {
		return
	}

	if _, err := os.Stat(appConfig.dataDir); err != nil { // most likely directory not exist
		if err = os.Mkdir(appConfig.dataDir, 0400); err != nil {
			// long line can be broken into multiple lines providing the last one is
			// '[', '{', '(', or an operator (as in this case)
			panic("Unable to access/create directory, " + appConfig.dataDir +
				", - " + err.Error())
		}
	}
	initilized = true
}

func constFilename(title string) string {
	return filepath.Join(appConfig.dataDir, title+".txt")
}

func (p *Page) save() error {
	return ioutil.WriteFile(constFilename(p.Title), p.Body, 0644)
}

func load(title string) (*Page, error) {
	body, err := ioutil.ReadFile(constFilename(title))
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

func loadPages() []Page {
	persistInit()

	var dataFileCompile = regexp.MustCompile("^[^/\\\\]+[/\\\\]([^/\\\\]+)[.]txt$") // TODO
	pages := []Page{}
	err := filepath.Walk(appConfig.dataDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("loadPages: path=", path)
		if info.IsDir() {
			if path == appConfig.dataDir {
				return nil
			}
			return filepath.SkipDir // sny subdirectory is ignored
		}
		m := dataFileCompile.FindStringSubmatch(path)
		if m != nil {
			p, e := load(m[1])
			if e != nil {
				fmt.Printf("unable to load %s - %s\n", m[1], e.Error())
			} else {
				pages = append(pages, *p)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Walk failed - ", err)
	}
	return pages
}
