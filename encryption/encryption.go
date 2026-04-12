package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
//	"fmt"
	"io"
//	"os"
//	"bufio"
//	"log"
	
)

func EncryptString(string string) ([]byte, []byte, error) {
	
	//make the byte information for both the text input as well as the random generated key
	text := []byte(string)
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, nil, err
	}
/*	err = os.WriteFile("keyFile.data", key, 0777)
	if err != nil {
		return nil, nil, err
	}
*/
	ciph, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	encryptedNonce := gcm.Seal(nonce, nonce, text, nil)


	return encryptedNonce, key, nil
}

//rework this so that it can return direct byte instead of string
func EncryptPassword (password string)(string, []byte, error) {

	encryptedPassword, key, err := EncryptString(password)
	if err != nil {
		return "", nil, err
	}

	encryptedPasswordString := string(encryptedPassword)

	return encryptedPasswordString, key, nil
}

func DecryptPassword (encryptedPassword string, key string) (string, error) {
	//placeholder function for decrypting password
	decryptedString, err := DecryptString(encryptedPassword, key)
	if err != nil {
		return "", err
	}

	return decryptedString, nil
}

func DecryptString (string string, key string) (string, error) {
	//placeholder function for decrypting string
}
