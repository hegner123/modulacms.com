package templates

import (
	"embed"
	"html/template"
)

//go:embed Components/*
var TemplateFiles embed.FS

func LoadTemplates(path string) (*template.Template, error) {
	var t *template.Template
	t, err := t.ParseFS(TemplateFiles)
	if err != nil {
		return nil, err
	}

	return t, nil
}
