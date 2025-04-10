package request

import "testing"

func TestGetContent(t *testing.T) {
	_, err := GetContent("/")
	if err != nil {
		t.Fatal(err)
	}

}
