package user

import (
	"bytes"
	"errors"
)

//Function CompaarePasswords needs to pull a databaseHash and not rely on passing the databaseHash, it should do this with a username string
func ComparePasswords(username string, newHash []byte) error {

	//this is a temp area to creat the base of the logic this will be removed in the futur for actual functionality
	databaseHash := []byte("Test" + "123" + "123")

	if len(newHash) == 0 {
		return errors.New("password cannot be empty string")
	}

	if !bytes.Equal(newHash, databaseHash){
		return errors.New("passwords are not the same")
	}

	return nil
}

//standin function for actual hashing
func HashPassword(password string, salt []byte, pepper []byte) ([]byte, error) {
	
	inputBytes := append([]byte(password), salt...)
	inputBytes = append(inputBytes, pepper...)

	return inputBytes, nil
}
