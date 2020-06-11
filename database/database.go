package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

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
	if _, err := os.Stat(viper.GetString("database_file")); os.IsNotExist(err) {
		return nil, errors.New("Database doesn't exists make sure you migrate the database first\n")
	}

	conn, err := sqlx.Connect("sqlite3", viper.GetString("database_file"))
	if err != nil {
		return nil, err
	}

	return &Database{conn}, nil
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
