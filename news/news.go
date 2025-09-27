package news

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type News struct {
	Title       string         `xml:"title"`
	Link        string         `xml:"link"`
	GUID        string         `xml:"guid"`
	PubDate     string         `xml:"pubDate"`
	Media       MediaThumbnail `xml:"media:thumbnail"`
	Description string         `xml:"description"`
}

type MediaThumbnail struct {
	// Captures the URL attribute: <media:thumbnail url="..."/>
	URL string `xml:"url,attr"`
	// Captures the width attribute: <media:thumbnail width="..."/>
	Width string `xml:"width,attr"`
	// Captures the height attribute: <media:thumbnail height="..."/>
	Height string `xml:"height,attr"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Image string `xml:"image"`
	Link  string `xml:"link"`
	Item  []News `xml:"item"`
}

func FetchNews(url string) {
	// Simulate fetching the URL
	resp, err := http.Get(url)

	if err != nil {
		println("Error fetching URL:", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("Error: received non-200 response code:", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Error reading response body:", err.Error())
		os.Exit(1)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)

	fmt.Printf("rss : %v\n", rss)
	if err != nil {
		println("Error parsing XML:", err.Error())
		os.Exit(1)
	}

	for _, news := range rss.Channel.Item {
		fmt.Printf("Title: %s\n Image : %s\nLink: %s\nDescription: %s\n\n", news.Title, news.Media.URL, news.Link, news.Description)
	}
}
