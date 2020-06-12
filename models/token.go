package models

// Token ..
type Token struct {
	ID     int64  `db:"id"`
	Token  string `db:"token"`
	UserID int64  `db:"user_id"`
}
