package crawler

import (
	"crypto/tls"
	"net/http"
)

type HttpFetcher struct {
	body string
	urls []string
}

func (f HttpFetcher) Fetch(url string) (string, []string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		return "", nil, err
	}
	b, u, err := ParseHTML(res.Body)
	if err != nil {
		return "", nil, err
	}
	return b, u, nil
}
