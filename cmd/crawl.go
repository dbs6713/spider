package cmd

import (
	"github.com/donbstringham/spider/src/crawler"
	"github.com/donbstringham/spider/src/models"
	"github.com/donbstringham/spider/src/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/spf13/jwalterweatherman"
)

func init() {
	RootCmd.AddCommand(crawlCmd)
	log.SetStdoutThreshold(log.LevelDebug)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Start the spider crawling",
	Long:  "Start the spider crawling with a set configuration or defaults",
	Run: func(cmd *cobra.Command, args []string) {
		var fetchCnt = 0
		var fetchTot = viper.GetInt("core.count")
		var seedURL = viper.GetString("core.seed")
		URLque := make([]string, 0)

		repo, err := storage.GetPageRepository("mem")
		if err != nil {
			log.FATAL.Fatal(err)
		}

		URLque = append(URLque, seedURL)

		for i := 0; fetchCnt < fetchTot || len(URLque) == 0; i++ {
			p, err := gettheweb(URLque[0])
			if err != nil {
				log.CRITICAL.Print(err)
			}

			// Add to queue and found URL's
			for x := 0; x < len(p.Urls); x++ {
				URLque = append(URLque, p.Urls[x])
			}

			// Remove the URL just fetched
			URLque = URLque[1:]

			err = repo.Add(p)
			if err != nil {
				log.CRITICAL.Print(err)
			}

			fetchCnt++
			log.INFO.Printf("Pages fetched: %d", fetchCnt)

			//log.INFO.Printf("URL: %s\nContent: %s\n",p.RawUrl, p.RawBody)
			log.INFO.Printf("URL: %s\n", p.RawUrl)
			for x := 0; x < len(p.Urls); x++ {
				log.INFO.Printf("     %s\n", p.Urls[x])
			}
		}

		c, err := repo.Count()
		if err != nil {
			log.CRITICAL.Print(err)
		}
		log.INFO.Printf("Pages fetched: %d", c)
	},
}

func gettheweb(u string) (*models.Page, error) {
	f := crawler.HttpFetcher{}
	b, urls, err := f.Fetch(u)
	if err != nil {
		return nil, err
	}

	p := models.NewPage(u)
	p.Fetched = true
	p.RawBody = b
	p.Urls = urls

	return p, err
}
