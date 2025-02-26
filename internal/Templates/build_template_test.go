package templates

import (
	"fmt"
	"testing"
)

func TestLoadTemplate(t *testing.T) {
	f, err := LoadTemplates()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(f.DefinedTemplates())

}
