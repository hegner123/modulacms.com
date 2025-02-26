package templates

import (
	"embed"
	"html/template"
)

//go:embed Components/*
var templateFiles embed.FS

func LoadTemplates() (*template.Template, error) {
	var t *template.Template
	t, err := t.ParseFS(templateFiles, "Components/*")
	if err != nil {
		return nil, err
	}

	return t, nil
}
