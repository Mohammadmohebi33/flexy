package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	FilePath string `koanf:"file_path"`
}

type SQLiteDB struct {
	config Config
	db     *sql.DB
}

func (s *SQLiteDB) Conn() *sql.DB {
	return s.db
}

func New(config Config) *SQLiteDB {

	db, err := sql.Open("sqlite3", config.FilePath)
	if err != nil {
		panic(fmt.Errorf("can't open sqlite db: %v", err))
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Minute * 3)

	return &SQLiteDB{config: config, db: db}
}
