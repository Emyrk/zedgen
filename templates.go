package zedgen

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
	"text/template"

	_ "embed"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed templates/*.tmpl
var templatesDirectory embed.FS

type Templates struct {
	All *template.Template
}

func LoadTemplates() (*Templates, error) {
	tpl := template.New("")

	err := fs.WalkDir(templatesDirectory, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walk templates: %w", err)
		}
		if d.IsDir() {
			return nil
		}

		tpl, err = tpl.ParseFS(templatesDirectory, path)
		if err != nil {
			return fmt.Errorf("parse templates: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("load templates: %w", err)
	}

	return &Templates{All: tpl}, nil
}

func capitalize(name string) string {
	return cases.Title(language.English).String(name)
}

func comment(body string) string {
	return "// " + strings.Join(strings.Split(body, "\n"), "\n// ")
}
