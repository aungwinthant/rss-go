package models

import (
	"os"
	"rss-go/db"
	"strings"
)

type News struct {
	Title       string
	Link        string
	GUID        string
	PubDate     string
	SiteID      int
	Image       string
	Description string
}

func (news *News) Save() (int, error) {

	m, _ := db.NewDBManager()
	defer m.Close()

	println("Saving news item:", news.Title)

	slug := strings.ToLower(strings.ReplaceAll(news.Title, " ", "-"))

	exists, err := m.DB.Query("SELECT COUNT(*) FROM news WHERE uuid = ?", news.GUID)
	if err != nil {
		println("Error checking for existing news item:", err.Error())
		return 0, err
	}
	if exists.Next() {
		println("News item already exists, skipping:", news.Title)
		return 0, err
	}
	res, err := m.DB.Exec("INSERT INTO news (uuid, title, title_slug, image, pub_date, link, description) VALUES (?, ?, ?, ?, ?, ?, ?)",
		news.GUID, news.Title, slug, news.Image, news.PubDate, news.Link, news.Description)

	if err != nil {
		println("Error saving news item:", err.Error())
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		println("Error getting last insert id:", err.Error())
		return 0, err
	}

	return int(lastID), nil

}
func (news News) ShowNews() {
	newsList := GetNews()

	if len(newsList) == 0 {
		println("No news items found. Please run go run main.go <sites.txt> to fetch news.")
		os.Exit(1)
	}

	println("Latest News:")
	for _, item := range newsList {
		println("Title:", item.Title)
		println("Link:", item.Link)
		println("PubDate:", item.PubDate)
		println("Description:", item.Description)
		println("Image URL:", item.Image)
		println("-----")
	}
}

func GetNews() []News {
	m, _ := db.NewDBManager()
	defer m.Close()

	rows, err := m.DB.Query("SELECT uuid, title, title_slug, image, pub_date, link, description FROM news LIMIT 10")
	if err != nil {
		println("Error querying news items:", err.Error())
		return nil
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var item News
		var slug interface{} // Placeholder for title_slug
		if err := rows.Scan(&item.GUID, &item.Title, &slug, &item.Image, &item.PubDate, &item.Link, &item.Description); err != nil {
			println("Error scanning news item:", err.Error())
			continue
		}
		newsList = append(newsList, item)
	}

	return newsList
}
