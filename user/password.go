package user

import (
	"bytes"
	"errors"
	"crypto/sha3"

	"golang.org/x/crypto/argon2"
)

//ComparePasswords needs to pull a databaseHash and not rely on passing the databaseHash, it should do this with a username string
func ComparePasswords(username string, newHash []byte) error {

	//this is a temp area to creat the base of the logic this will be removed in the future for actual functionality
	databaseHash, _ := HashPassword("Test", []byte("123"), []byte("123"))


	if len(newHash) == 0 {
		return errors.New("password cannot be empty string")
	}

	if !bytes.Equal(newHash, databaseHash){
		return errors.New("passwords are not the same")
	}

	return nil
}

//HashPassword creates 64 byte masterHash that can be split to 2 32 byte hashes for other things might change later depending what I learn
func HashPassword(password string, salt []byte, pepper []byte) ([]byte, error) {
	h := sha3.New256()
	h.Write(pepper)
	h.Write([]byte(password))
	pepperedPassword := h.Sum(nil)

	masterHash := argon2.IDKey(pepperedPassword, salt, 3, 64*1024, 4, 64)

	return masterHash, nil
}
