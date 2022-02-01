package pet

import "time"

type Pet struct {
	ID        int       `db:"id"`
	Animal    string    `db:"animal"`
	Breed     string    `db:"breed"`
	Price     float64   `db:"price"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
