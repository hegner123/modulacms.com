package models

import (
	"encoding/json"
	"fmt"
)

func BuildRoot(data []byte) (Root, error) {
	r := Root{}

	err := json.Unmarshal(data, &r)
    if err!=nil {
        fmt.Println(err)
    }

	return r, nil
}
