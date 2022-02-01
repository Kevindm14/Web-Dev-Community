package sale

import (
	"petstore/manage/client"
	"petstore/manage/pet"
)

type Sale struct {
	Client client.Client
	Pets   []pet.Pet
}
