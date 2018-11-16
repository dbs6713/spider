package src

import (
	"net/http"
)

type HttpFetcher struct {
	body string
	urls []string
}

func (f HttpFetcher) Fetch(url string) (string, []string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}

	b, u, err := ParseHTML(res.Body)
	if err != nil {
		return "", nil, err
	}

	return b, u, nil
}
