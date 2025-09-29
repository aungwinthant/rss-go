package main

import (
	"bufio"
	"os"

	"rss-go/internal/models"
)

func main() {
	args := os.Args[1:]

	var news models.News
	if len(args) < 1 {
		news.ShowNews()
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
		news := fetchrss.Fetch(url)
		println("Successfully fetched URL:", url)
		SaveNewsList(news)
	}

	if err := scanner.Err(); err != nil {
		println("Error reading file:", err.Error())
	}

	if err != nil {
		println("Error fetching URL:", err.Error())
		os.Exit(1)
	}
}

func SaveNewsList(news []models.News) {

	for _, item := range news {
		println("Saving news item:", item.Title)
		item.Save()
	}
}
