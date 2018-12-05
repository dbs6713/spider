package models

import (
	"encoding/json"
	"errors"
)

type Page struct {
	Fetched bool     `json:"fetched,omitempty"`
	RawBody string   `json:"body,omitempty"`
	RawUrl  string   `json:"url,omitempty"`
	Urls    []string `json:"urls,omitempty"`
}

func NewPage(u string) *Page {
	return &Page{Fetched: false, RawUrl: u}
}

func (p *Page) ToJSON() (string, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return "", errors.New("ERROR: JSON conversion")
	}
	return string(bs), nil
}

func (p *Page) FromJSON(JSON []byte) (*Page, error) {
	np := &Page{}
	if err := json.Unmarshal(JSON, np); err != nil {
		return nil, err
	}
	return np, nil
}
