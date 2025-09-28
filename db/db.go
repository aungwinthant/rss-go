package db

import (
	"database/sql"
	"fmt"
	"os"

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

	return &DBManager{DB: db}, nil
}

func InitDB() *DBManager {
	m, err := NewDBManager()
	if err != nil {
		println("Error initializing DBManager:", err.Error())
		os.Exit(1)
	}
	err = m.CreateTables()
	if err != nil {
		println("Error creating tables:", err.Error())
		os.Exit(1)
	}
	return m
}

func (m *DBManager) Close() {
	m.DB.Close()
}

func (m *DBManager) CreateTables() error {
	// Create a table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS news (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"uuid" TEXT,		
		"title" TEXT,
		"title_slug" TEXT,
		"image" TEXT,
		"pub_date" TEXT,	
		"link" TEXT,
		"description" TEXT
	  );`

	_, err := m.DB.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	return nil
}
