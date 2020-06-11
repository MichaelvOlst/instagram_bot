package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" //sqlite3 driver
	"github.com/spf13/viper"
)

// Database holds the connection
type Database struct {
	*sqlx.DB
}

// New open and migrating the database
func New() (*Database, error) {

	// TODO check if file exists for migrating

	conn, err := sqlx.Connect("sqlite3", viper.Get("database_file").(string))
	if err != nil {
		return nil, err
	}

	return &Database{conn}, nil
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
