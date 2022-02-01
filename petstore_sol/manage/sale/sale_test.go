package sale_test

import (
	"log"
	"petstore/manage/client"
	"petstore/manage/pet"
	"petstore/manage/sale"
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

		c, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)
		r.NotZero(c)

		_, err = sale.DB(tx).Create(c.ID, p.ID)
		r.Nil(err)

		s, err := sale.DB(tx).ForClient(c.ID)
		r.Nil(err)
		r.NotNil(s)
	})
}

func Test_FindForClient(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		p, err := pet.DB(tx).Create("cat", "Abyssinian", 1000, 1)
		r.Nil(err)
		r.NotZero(p)

		c, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)
		r.NotZero(c)

		_, err = sale.DB(tx).Create(c.ID, p.ID)
		r.Nil(err)

		s, err := sale.DB(tx).ForClient(c.ID)
		r.Nil(err)
		r.Len(s.Pets, 1)
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
