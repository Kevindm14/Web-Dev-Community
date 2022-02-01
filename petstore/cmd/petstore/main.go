package main

import (
	"fmt"

	"github.com/gobuffalo/pop/v6"
)

func main() {
	db, err := pop.Connect("")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Conected to database! yay!")
}
