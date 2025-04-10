package templates

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hegner123/modulacms.com/internal/models"
)

func LoadTemplates() (*template.Template, error) {
	var t *template.Template
	t, err := t.ParseGlob("./Components/*")
	if err != nil {
		return nil, err
	}

	return t, nil
}

func BuildPage(root models.Root, w http.ResponseWriter, r *http.Request) error {
	t, err := LoadTemplates()
	if err != nil {
		return err
	}

    t.Funcs(template.FuncMap{
		"eq": eq,
	})

	err = t.ExecuteTemplate(w, "layout.html", root)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil

}
