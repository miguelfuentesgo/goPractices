package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Run() {

	urlObjective := "https://news.ycombinator.com/"

	// HTPP Petition

	res, err := http.Get(urlObjective)

	if err != nil {
		log.Fatal("ERROR doing HTTP request", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("ERROR response status code is %d", res.StatusCode)
	}

	//Use goquery to obtain HTML from the response

	doc, err := goquery.NewDocumentFromReader(res.Body)

	// Extract titles from document
	fmt.Println("Last news from Ycombinator")
	doc.Find(".titleline").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		log.Printf("Title %d: %s", i, title)
	})

}
