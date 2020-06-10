package database

import (
	"database/sql"
	"fmt"
	"instagram_bot/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" //sqlite3 driver
)

// Database holds the connection
type Database struct {
	*sqlx.DB
}

// New open and migrating the database
func New(cfg *config.Config) (*Database, error) {

	conn, err := sqlx.Connect("sqlite3", cfg.Database.Name)
	if err != nil {
		return nil, err
	}

	db := &Database{conn}

	db.migrate()

	err = db.createAdminUser(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Close the db pool
func (db *Database) Close() error {
	return db.Close()
}

func (db *Database) rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}
