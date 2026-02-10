package zedgen

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templatesDirectory embed.FS

type Templates struct {
	All *template.Template
}

func LoadTemplates() (*Templates, error) {
	var tpl = template.New("").Funcs(template.FuncMap{
		"capitalize": capitalize,
		"comment":    comment,
	})
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
	if len(name) == 0 {
		return name
	}
	// Only uppercase the first letter, preserve the rest
	return strings.ToUpper(name[:1]) + name[1:]
}

func comment(body string) string {
	return "// " + strings.Join(strings.Split(body, "\n"), "\n// ")
}
