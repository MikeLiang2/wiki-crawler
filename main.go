package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// Ensure output folders exist
	if err := os.MkdirAll("output/wikipages", os.ModePerm); err != nil {
		log.Fatal("Failed to create folders:", err)
	}

	// Load URLs from urls.json
	urls, err := loadURLs("urls.json")
	if err != nil {
		log.Fatal("Failed to load URLs:", err)
	}

	// Create output file
	jsonFile, err := os.Create("output/texts.jl")
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer jsonFile.Close()

	// Set up and start crawler
	c := setupCollector(jsonFile)
	for _, u := range urls {
		c.Visit(u)
	}
	c.Wait()

	fmt.Printf("Finished in %.2f seconds\n", time.Since(start).Seconds())
}
