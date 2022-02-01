package sale

import (
	"errors"
	"petstore/manage/client"
	"petstore/manage/pet"
	"time"

	"github.com/gobuffalo/pop/v6"
)

type Database interface {
	Create(clientID, petID int) (*Sale, error)
	ForClient(clientID int) (*Sale, error)
}

func DB(tx *pop.Connection) Database {
	return &db{tx}
}

type db struct {
	tx *pop.Connection
}

func (db *db) Create(clientID, petID int) (*Sale, error) {
	p, err := pet.DB(db.tx).Find(pet.WithID(petID))
	if err != nil {
		return nil, err
	}

	c, err := client.DB(db.tx).Find(client.WithID(clientID))
	if err != nil {
		return nil, err
	}

	err = db.tx.RawQuery("insert into sales values(?,?,?,?)", clientID, petID, time.Now(), time.Now()).Exec()
	if err != nil {
		return nil, err
	}

	s := &Sale{
		Client: *c[0],
		Pets:   []pet.Pet{*p[0]},
	}

	return s, nil
}

func (db *db) ForClient(clientID int) (*Sale, error) {
	c, err := client.DB(db.tx).Find(client.WithID(clientID))
	if err != nil {
		return nil, err
	}

	if len(c) == 0 {
		return nil, errors.New("client ID does not exist")
	}

	var pets []pet.Pet
	q := db.tx.Q()
	q.LeftJoin("sales s", "s.pet_id = pets.id").
		Where("s.client_id = ?", c[0].ID).All(&pets)

	return &Sale{*c[0], pets}, nil
}
