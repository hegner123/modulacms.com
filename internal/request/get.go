package request

import (
	"fmt"
	"io"
	"net/http"
)

func GetContent(path string) ([]byte, error) {
	resp, err := http.Get("http://localhost:4000/" + path)
	if err != nil {
		e := fmt.Errorf("error making get request: %v", err)
		return nil, e
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		e := fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		return nil, e
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		e := fmt.Errorf("error reading response body: %v", err)
		return nil, e
	}

	return body, nil
}
