package client

import (
	"errors"

	"github.com/gobuffalo/pop/v6"
)

type Database interface {
	List() ([]*Client, error)
	Find(...FindSpec) ([]*Client, error)
	Create(name, lastname, email, address, phone, gender string, age int) (*Client, error)
	Update(*Client) error
	Destroy(*Client) error
}

// DB returns a Database implementation to operate with
// pets database.
func DB(tx *pop.Connection) Database {
	return &db{tx}
}

type db struct {
	tx *pop.Connection
}

// Create creates the pet with values passed as arguments
func (db *db) Create(name, lastname, email, address, phone, gender string, age int) (*Client, error) {
	client := &Client{
		Name:     name,
		Lastname: lastname,
		Email:    email,
		Address:  address,
		Phone:    phone,
		Gender:   gender,
		Age:      age,
	}
	return client, db.tx.Create(client)
}

func (db *db) Update(c *Client) error {
	if c.ID == 0 {
		return errors.New("can not update an inexistent client")
	}
	return db.tx.Update(c)
}

func (db *db) List() ([]*Client, error) {
	var clients []*Client
	return clients, db.tx.Q().All(&clients)
}

func (db *db) Destroy(c *Client) error {
	if c.ID == 0 {
		return errors.New("can not destroy an inexistent client")
	}

	return db.tx.Destroy(c)
}

func (db *db) Find(specs ...FindSpec) ([]*Client, error) {
	q := db.tx.Q()
	for _, spec := range specs {
		spec(q)
	}

	var clients []*Client
	return clients, db.tx.All(&clients)
}

// FindSpec some filter specification to find clients.
type FindSpec func(q *pop.Query)

// WithID an specification to filter by id.
func WithID(id int) FindSpec {
	return func(q *pop.Query) {
		q.Where("id = ?", id)
	}
}

// WithFullName an specification to filter by name or lastname.
func WithFullName(fullName string) FindSpec {
	return func(q *pop.Query) {
		q.Where("name||' '||lastname ILIKE ?", "%"+fullName+"%")
	}
}

// WithEmail an specification to filter by email.
func WithEmail(email string) FindSpec {
	return func(q *pop.Query) {
		q.Where("email = ?", email)
	}
}
