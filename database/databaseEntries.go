package database

import (
	"fmt"
	encrypt "github.com/edickens09/passwordVault/encryption"
)

//rework this so that you can return the byte directly instead of converting to string. Need to refactor other functions first
func EncryptPassword(password string)(string, []byte, error) {
	
	passwordHash, key, err := encrypt.EncryptString(password)
	if err != nil {
		return "", nil, err
	}
	passwordHashString := string(passwordHash)
	fmt.Println("Database Encrypt Password is working " + passwordHashString)

	return passwordHashString, key, nil
}
