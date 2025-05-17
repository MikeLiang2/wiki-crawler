package main

import (
	"encoding/json"
	"os"
	"strings"
)

// Load URLs from a JSON file
func loadURLs(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		return nil, err
	}
	return urls, nil
}

// Save raw HTML to file
func saveHTML(title string, body []byte) {
	filename := "output/wikipages/" + sanitizeFilename(title) + ".html"
	if err := os.WriteFile(filename, body, 0644); err != nil {
		println("Failed to save HTML:", err.Error())
	}
}

// Sanitize filename by replacing spaces and slashes
func sanitizeFilename(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "-")
	return name
}
