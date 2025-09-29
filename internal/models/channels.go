package models

import (
	"database/sql"
	"fmt"
	"rss-go/db"
)

type Channel struct {
	Title string
	Image string
	Link  string
	Item  []News
}

func (channel *Channel) SaveChannel(name string) (int, error) {
	m, err := db.NewDBManager()
	if err != nil {
		return 0, err
	}
	defer m.Close()

	var channelID int
	err = m.DB.QueryRow("SELECT id FROM channels WHERE name = ?", name).Scan(&channelID)
	if err == sql.ErrNoRows {
		res, err := m.DB.Exec("INSERT INTO channels (name) VALUES (?)", name)
		if err != nil {
			return 0, fmt.Errorf("error inserting channel: %w", err)
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("error getting last insert id: %w", err)
		}
		channelID = int(lastID)
	} else if err != nil {
		return 0, fmt.Errorf("error querying channel: %w", err)
	}
	return channelID, nil
}
