package pointers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Ex6 -Exported function to main - bcyypt
func Ex6() {

	password := "Hello world"

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("\n Correct Password!!")
	}
}
