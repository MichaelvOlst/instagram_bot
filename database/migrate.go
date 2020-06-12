package database

// Migrate migrates the database
func (db *Database) Migrate() {

	var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tokens (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		token VARCHAR(255) NOT NULL
	);
	FOREIGN KEY (user_id) REFERENCES users (id);
	`
	db.Exec(schema)

}
