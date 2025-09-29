package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DBManager struct {
	DB *sql.DB
}

func NewDBManager() (*DBManager, error) {
	// Initialize your database connection here
	db, err := sql.Open("sqlite3", "./news.db")
	if err != nil {
		println("Error opening database:", err.Error())
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	// 2. Ping to verify the connection is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := CreateTables(db); err != nil {
		db.Close() // IMPORTANT: Close connection on failure
		return nil, fmt.Errorf("error setting up database schema: %w", err)
	}

	return &DBManager{DB: db}, nil
}

func (m *DBManager) Close() {
	m.DB.Close()
}

func CreateTables(db *sql.DB) error {
	// Create a table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Channel (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"link" TEXT,
		"image" TEXT,)
	);
	
	CREATE TABLE IF NOT EXISTS news (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"uuid" TEXT,		
		"title" TEXT,
		"title_slug" TEXT,
		'channel_id' INTEGER,
		"image" TEXT,
		"pub_date" TEXT,	
		"link" TEXT,
		"description" TEXT
	  );

	CREATE UNIQUE INDEX IF NOT EXISTS idx_news_uuid ON news (uuid);
	CREATE INDEX IF NOT EXISTS idx_news_channel_id ON news (channel_id);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating tables : %w", err)
	}
	return nil
}
