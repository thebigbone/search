package main

import (
	"fmt"
	"regexp"

	colly "github.com/gocolly/colly"
	"golang.org/x/exp/rand"
)

type ArchiveScraper struct {
	c *colly.Collector
}

func (s *ArchiveScraper) handleArchiveLink(e *colly.HTMLElement) {
	link := e.Attr("href")
	if patternMatch(link) {
		fmt.Printf("Archive link found: %q -> %s\n", e.Text, link)
		s.c.Visit(e.Request.AbsoluteURL(link))
	}
}

func (s *ArchiveScraper) startScraping() {
	s.c.OnHTML("a[href]", s.handleArchiveLink)

	s.c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	s.c.Visit(url)
}

func patternMatch(link string) bool {
	re := regexp.MustCompile(`(\d{4}/schedule)`)
	match := re.FindStringSubmatch(link)

	if len(match) > 0 {
		return true
	} else {
		return false
	}
}

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
