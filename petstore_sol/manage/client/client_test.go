package client_test

import (
	"log"
	"petstore/manage/client"
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
		c, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)
		r.NotZero(c)
	})
}

func Test_List(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		_, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)

		clients, err := client.DB(tx).List()
		r.Nil(err)
		r.Len(clients, 1)
	})
}

func Test_Find(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)

		_, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)

		clients, err := client.DB(tx).Find(
			client.WithEmail("paguilar@test.com.co"),
			client.WithFullName("PepO"),
		)
		r.Nil(err)
		r.Len(clients, 1)
	})
}

func Test_Destroy(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		c, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)

		r.Nil(client.DB(tx).Destroy(c))
		clients, err := client.DB(tx).List()
		r.Nil(err)
		r.Len(clients, 0)
	})
}

func Test_Update(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)
		c, err := client.DB(tx).Create("pepo", "aguilar", "paguilar@test.com.co", "test address", "test phone", "Male", 15)
		r.Nil(err)

		c.Lastname = "Aguilar"
		r.Nil(client.DB(tx).Update(c))

		clients, err := client.DB(tx).Find()
		r.Nil(err)
		r.Len(clients, 1)
		r.Equal("Aguilar", clients[0].Lastname)
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
