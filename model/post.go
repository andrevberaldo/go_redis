package model

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) ToJSON() ([]byte, error) {
	json, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Enable to convert Product to JSON")
		return []byte{1}, err
	}

	return json, nil
}
