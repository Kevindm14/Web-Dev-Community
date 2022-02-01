package client

import "time"

type Client struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Lastname  string    `db:"lastname"`
	Email     string    `db:"email"`
	Address   string    `db:"address"`
	Phone     string    `db:"phone"`
	Gender    string    `db:"gender"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
