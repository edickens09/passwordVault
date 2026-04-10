package user

import (
	"errors"
)

func ComparePasswords(newHash string, databaseHash string) error {

	if newHash == "" {
		return errors.New("password cannot be empty string")
	}

	if newHash != databaseHash {
		return errors.New("passwords are not the same")
	}

	return nil
}

func HashPassword(password string) (string, error) {
	
	passwordHash := password+"123"

	return passwordHash, nil
}
