package crawler

type Crawler interface {
	Crawl(u string) string
}
