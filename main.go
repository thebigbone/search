package main

import (
	colly "github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const url = "https://recon.cx/2024/archive.html"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("cfp.recon.cx", "recon.cx"),
	)
	scraper := &ArchiveScraper{c: c}
	scraper.startScraping()
}
