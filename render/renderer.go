package render

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Renderer struct {
	*template.Template
}

func New() *Renderer {
	tmpl := template.New("")
	err := filepath.Walk("templates/", func(path string, _ os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			if _, err = tmpl.ParseFiles(path); err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return &Renderer{tmpl}
}
