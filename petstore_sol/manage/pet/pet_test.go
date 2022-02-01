package pet_test

import (
	"log"
	"petstore/manage/pet"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		p, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)
		r.NotZero(p)
	})
}

func Test_List(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		_, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)

		_, err = pet.DB(tx).Create("dog", "Pitbull", 2000, 1)
		r.Nil(err)

		pets, err := pet.DB(tx).List()
		r.Nil(err)
		r.Len(pets, 2)
	})
}

func Test_Find(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)

		_, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)

		pets, err := pet.DB(tx).Find(
			pet.WithAge(1),
			pet.WithAnimal("cat"),
			pet.WithPriceLessThan(1500),
		)
		r.Nil(err)
		r.Len(pets, 1)
	})
}

func Test_Destroy(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		p, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)

		r.Nil(pet.DB(tx).Destroy(p))
		pets, err := pet.DB(tx).List()
		r.Nil(err)
		r.Len(pets, 0)
	})
}

func Test_Update(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		p, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)

		p.Breed = "Ragdoll"
		r.Nil(pet.DB(tx).Update(p))

		pets, err := pet.DB(tx).Find(pet.WithAnimal("cat"))
		r.Nil(err)
		r.Len(pets, 1)
		r.Equal("Ragdoll", pets[0].Breed)
	})
}

var db *pop.Connection

func init() {
	pop.Debug = true
	pop.AddLookupPaths("../../")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	var err error
	db, err = pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}
}

func transaction(fn func(tx *pop.Connection)) {
	err := db.Rollback(func(tx *pop.Connection) {
		fn(tx)
	})
	if err != nil {
		log.Fatal(err)
	}
}
