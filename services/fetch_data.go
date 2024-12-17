package services

import (
	"fmt"
	"io"
	"net/http"
)

func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Unable to fetch data")
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Unable to parse response body")
		return nil, err
	}

	return body, nil
}
