package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type PageData struct {
	Title string   `json:"title"`
	URL   string   `json:"url"`
	Text  string   `json:"text"`
	Tags  []string `json:"tags"`
}

// Set up Colly collector with all handlers
func setupCollector(jsonFile *os.File) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
		Delay:       1 * time.Second,
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		handlePage(e, jsonFile)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	return c
}

// Handle scraping and writing logic
func handlePage(e *colly.HTMLElement, jsonFile *os.File) {
	title := e.DOM.Find("#firstHeading").Text()
	url := e.Request.URL.String()
	text := strings.TrimSpace(e.DOM.Find("div.mw-parser-output").Text())

	var tags []string
	e.DOM.Find("#mw-normal-catlinks ul li a").Each(func(_ int, s *goquery.Selection) {
		tag := strings.TrimSpace(s.Text())
		if tag != "" {
			tags = append(tags, tag)
		}
	})

	page := PageData{
		Title: title,
		URL:   url,
		Text:  text,
		Tags:  tags,
	}

	data, err := json.Marshal(page)
	if err != nil {
		log.Printf("Failed to marshal %s: %v", title, err)
		return
	}
	jsonFile.Write(data)
	jsonFile.Write([]byte("\n"))

	saveHTML(title, e.Response.Body)
	log.Println("Scraped:", title)
}
