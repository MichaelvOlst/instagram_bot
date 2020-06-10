package database

func (db *Database) migrate() {

	var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);
	`
	db.MustExec(schema)

}
