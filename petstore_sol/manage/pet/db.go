package pet

import (
	"errors"

	"github.com/gobuffalo/pop/v6"
)

type Database interface {
	List() ([]*Pet, error)
	Find(...FindSpec) ([]*Pet, error)
	Create(animal, breed string, price float64, age int) (*Pet, error)
	Update(*Pet) error
	Destroy(*Pet) error
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
func (db *db) Create(animal, breed string, price float64, age int) (*Pet, error) {
	pet := &Pet{
		Animal: animal,
		Breed:  breed,
		Price:  price,
		Age:    age,
	}
	return pet, db.tx.Create(pet)
}

func (db *db) Update(p *Pet) error {
	if p.ID == 0 {
		return errors.New("can not update an inexistent pet")
	}
	return db.tx.Update(p)
}

func (db *db) List() ([]*Pet, error) {
	var pets []*Pet
	return pets, db.tx.Q().All(&pets)
}

func (db *db) Destroy(p *Pet) error {
	if p.ID == 0 {
		return errors.New("can not destroy an inexistent pet")
	}

	return db.tx.Destroy(p)
}

func (db *db) Find(specs ...FindSpec) ([]*Pet, error) {
	q := db.tx.Q()
	for _, spec := range specs {
		spec(q)
	}

	var pets []*Pet
	return pets, db.tx.All(&pets)
}

// FindSpec some filter specification to find pets.
type FindSpec func(q *pop.Query)

// WithID an specification to filter by id.
func WithID(id int) FindSpec {
	return func(q *pop.Query) {
		q.Where("id = ?", id)
	}
}

// WithAnimal an specification to filter by animal.
func WithAnimal(animal string) FindSpec {
	return func(q *pop.Query) {
		q.Where("animal = ?", animal)
	}
}

// WithPriceLessThan an specification to filter by price.
func WithPriceLessThan(price float64) FindSpec {
	return func(q *pop.Query) {
		q.Where("price <= ?", price)
	}
}

// WithAge an specification to filter by age.
func WithAge(age int) FindSpec {
	return func(q *pop.Query) {
		q.Where("age = ?", age)
	}
}
