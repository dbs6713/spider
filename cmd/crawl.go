package cmd

import (
	"github.com/donbstringham/spider/src/crawler"
	"github.com/donbstringham/spider/src/models"
	"github.com/donbstringham/spider/src/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"

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
		var user = viper.GetString("db.user")
		var pass = viper.GetString("db.pass")
		var dbn = viper.GetString("db.dbname")
		var host = viper.GetString("db.host")
		var port = viper.GetString("db.port")

		URLque := make([]string, 0)
		ch := make(chan *models.Page)

		start := time.Now()

		//repo, err := storage.GetPageRepository("mem")
		repo, err := storage.GetPageRepository("mysql", user, pass, host, port, dbn)
		if err != nil {
			log.FATAL.Fatal(err)
		}

		fetcher := crawler.NewHttpFetcher()

		URLque = append(URLque, seedURL)

		for i := 0; fetchCnt < fetchTot || len(URLque) == 0; i++ {
			go makeRequest(URLque[0], fetcher, ch)

			p := <-ch

			if p == nil {
				log.ERROR.Printf("ERROR: page fetch %s", URLque[0])
				continue
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

		secs := time.Since(start).Seconds()

		log.INFO.Printf("%.2f elapsed", secs)
	},
}

func makeRequest(u string, f crawler.Fetcher, ch chan<- *models.Page) {
	b, urls, err := f.Fetch(u)
	if err != nil {
		log.ERROR.Println(err)
		ch <- nil
		return
	}

	p := models.NewPage(u)
	p.Fetched = true
	p.RawBody = b
	p.Urls = urls

	ch <- p
}
