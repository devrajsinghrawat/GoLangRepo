package exsoll7

import (
	"encoding/json"
	"fmt"
	"os"
)

// Ex1

type user struct {
	Name   string
	Last   string
	Age    int
	Saying []string
}

// Ex1 -
func Ex1() {

	u1 := user{"dev", "raj", 30, []string{"hello world", "Bolo wolrd"}}
	u2 := user{"raj", "dev", 35, []string{"hello world 1", "Bolo wolrd 1"}}

	users := []user{u1, u2}

	// pointer to Encoder should call encode

	// func (enc *Encoder)

	// encoderPointer := json.NewEncoder(os.Stdout)
	// err := encoderPointer.Encode(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("encoding done \n")
	}
}
