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

	if err != nil {
		println("Error opening file:", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		news := news.FetchNews(url)
		println("Successfully fetched URL:", url)
		SaveNewsList(news)
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

func SaveNewsList(news []news.News) {
	// Save the news item to the database
	// This is a placeholder function; implement actual DB logic here
	for _, item := range news {
		println("Saving news item:", item.Title)
		item.Save()
	}
}
