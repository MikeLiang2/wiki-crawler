# Wikipedia Web Crawler (Go + Colly)

This project implements a concurrent web crawler in Go using the Colly framework. It extracts textual content and category tags from a list of Wikipedia pages related to intelligent systems and robotics. The result is saved in a JSON Lines (.jl) format, suitable for input to a knowledge base system.

---

## Features

- Concurrent scraping using Colly's async mode and goroutines
- Extracts page title, plain text content, and category tags
- Saves raw HTML files for each Wikipedia page
- Outputs a structured `.jl` file for further processing

---

## Prerequisites

- Go 1.17 or later
- Colly and Goquery libraries (installed via `go mod tidy`)

---

## How to Run

Clone the repository and run the program using:
go run .

The input file urls.json should contain a list of Wikipedia URLs in JSON array format:
## Example urls.json

[
  "https://en.wikipedia.org/wiki/Robotics",
  "https://en.wikipedia.org/wiki/Robot",
  "https://en.wikipedia.org/wiki/Chatbot"
]

## Output
The file output/texts.jl will contain one JSON object per line, representing one Wikipedia page. Example:

{
  "title": "Robotics",
  "url": "https://en.wikipedia.org/wiki/Robotics",
  "text": "Robotics is an interdisciplinary branch...",
  "tags": ["Automation", "Artificial intelligence"]
}

Raw HTML is saved to output/wikipages/<Title>.html.

## Performance Comparison
Language	Time (10 pages)
Python (Scrapy)	12.63 seconds
Go (Colly)	1.29 seconds (Sending 10 request a time)

Go's performance benefits from native concurrency and lower I/O overhead.

## Testing
This project includes a unit test for filename sanitization.

Run tests with:
go test

## GenAI Tools
This project used OpenAI ChatGPT-4 to assist with:
    Code structure design and modularization
    Fixing main.go
    Generating documentation and test code
Documentation of AI usage is provided as required:
