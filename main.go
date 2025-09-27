package main

import (
	"bufio"
	"os"
	"rss-go/news"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("Usage: go run main.go <sites.txt>")
		os.Exit(1)
	}

	filename := args[0]

	file, err := os.Open(filename)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		news.FetchNews(url)
		println("Successfully fetched URL:", url)
	}

	if err := scanner.Err(); err != nil {
		println("Error reading file:", err.Error())
	}

	// _, err := fetchURL(url)
	if err != nil {
		println("Error fetching URL:", err.Error())
		os.Exit(1)
	}
}
