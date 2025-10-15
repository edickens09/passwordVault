package database

import (
	"fmt"
)

func EncryptPassword(password string)(string, error) {
	
	passwordHash := password + "123"
	fmt.Println("Database Encrypt Password is working " + passwordHash)

	return passwordHash, nil
}
