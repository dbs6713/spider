package crawltypes

type CrawlHttp struct {
	Crawled    bool
	RawContent string
	RawUrl     string
}

func NewCrawlHttp() *CrawlHttp {
	return &CrawlHttp{
		Crawled:    false,
		RawContent: "",
		RawUrl:     "",
	}
}

func (ch *CrawlHttp) Crawl(u string) string {
	return ""
}
