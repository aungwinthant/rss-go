package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("Usage: go run main.go <URL>")
		os.Exit(1)
	}
	url := args[0]

	_, err := fetchURL(url)
	if err != nil {
		println("Error fetching URL:", err.Error())
		os.Exit(1)
	}

	println("Successfully fetched URL:", url)

}

func fetchURL(url string) (string, error) {
	// Simulate fetching the URL
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	fmt.Printf("Response Body : %v\n\n", rss)

	// In a real scenario, you would read and return the response body here.
	// For simplicity, we just return a success message.
	return "Fetched content from " + url, nil
}
