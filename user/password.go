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

//standin function for actual hashing
func HashPassword(password string, salt string) (string, error) {
	
	passwordHash := password+salt

	return passwordHash, nil
}
