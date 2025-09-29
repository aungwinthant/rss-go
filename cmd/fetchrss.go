package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type XMLNews struct {
	Title       string            `xml:"title"`
	Link        string            `xml:"link"`
	GUID        string            `xml:"guid"`
	PubDate     string            `xml:"pubDate"`
	Media       XMLMediaThumbnail `xml:"media:thumbnail"`
	Description string            `xml:"description"`
}

type XMLMediaThumbnail struct {
	// Captures the URL attribute: <media:thumbnail url="..."/>
	URL string `xml:"url,attr"`
	// Captures the width attribute: <media:thumbnail width="..."/>
	Width string `xml:"width,attr"`
	// Captures the height attribute: <media:thumbnail height="..."/>
	Height string `xml:"height,attr"`
}

type RSS struct {
	XMLName xml.Name   `xml:"rss"`
	Channel XMLChannel `xml:"channel"`
}

type XMLChannel struct {
	Title string    `xml:"title"`
	Image string    `xml:"image"`
	Link  string    `xml:"link"`
	Item  []XMLNews `xml:"item"`
}

func init() {
	// Initialization code, if needed
}
func Fetch(url string) RSS {
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

	return rss
}
