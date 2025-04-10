package models

import (
	"fmt"
	"os"
	"testing"
)

func TestUnmarshallData(t *testing.T) {
	f, err := os.ReadFile("content.json")
	if err != nil {
		t.Fatal(err)
	}
	m, err := BuildRoot(f)
    if err!=nil {
		t.Fatal(err)
    }
	fmt.Println(m.Render())

}
