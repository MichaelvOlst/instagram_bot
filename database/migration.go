package database

// Migrate migrates the database
func (db *Database) Migrate() {

	var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	CREATE UNIQUE INDEX email ON users (email);
	`
	db.MustExec(schema)

}
