package database

import (
	"fmt"
	encrypt "github.com/edickens09/passwordVault/encryption"
)

func EncryptPassword(password string)(string, error) {
	
	passwordHash := encrypt.EncryptString(password)
	fmt.Println("Database Encrypt Password is working " + passwordHash)

	return passwordHash, nil
}
