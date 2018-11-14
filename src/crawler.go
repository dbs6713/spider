package src

type Crawler interface {
	Crawl(u string) string
}
