package crawler

import (
	"crypto/tls"
	"net/http"
)

type HttpFetcher struct {
	body   string
	client *http.Client
	urls   []string
}

func NewHttpFetcher() *HttpFetcher {
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: t}
	return &HttpFetcher{client: c}
}

func (f HttpFetcher) Fetch(url string) (string, []string, error) {
	res, err := f.client.Get(url)
	if err != nil {
		return "", nil, err
	}
	b, u, err := ParseHTML(res.Body)
	if err != nil {
		return "", nil, err
	}
	return b, u, nil
}
