package user

import (
	"fmt"
	"errors"
)

func ComparePasswords(password string) error {

	if password == "" {
		return errors.New("password cannot be empty string")
	}

	fmt.Println("This looks ok so far: " + password)
	return nil
}

func HashPassword(password string) (string, error) {
	
	passwordHash := password+"123"
	fmt.Println("HashPassword Function call is working " + passwordHash)

	return passwordHash, nil
}
