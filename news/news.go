package news

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"rss-go/db"
	"strings"
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

func FetchNews(url string) []News {
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

	return rss.Channel.Item
}

func (news *News) Save() {

	m := db.InitDB()
	// Save the news item to the database
	// This is a placeholder function; implement actual DB logic here
	println("Saving news item:", news.Title)

	slug := strings.ToLower(strings.ReplaceAll(news.Title, " ", "-"))

	exists, err := m.DB.Query("SELECT COUNT(*) FROM news WHERE uuid = ?", news.GUID)
	if err != nil {
		println("Error checking for existing news item:", err.Error())
		return
	}
	if exists.Next() {
		println("News item already exists, skipping:", news.Title)
		return
	}
	_, err = m.DB.Exec("INSERT INTO news (uuid, title, title_slug, image, pub_date, link, description) VALUES (?, ?, ?, ?, ?, ?, ?)",
		news.GUID, news.Title, slug, news.Media.URL, news.PubDate, news.Link, news.Description)

	if err != nil {
		println("Error saving news item:", err.Error())
	}

	defer m.Close()
}
