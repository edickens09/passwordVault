package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("Decryption Program v.0.00001")

	logFile, err := os.OpenFile("decryptionLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error with log file")
	}

	log.SetOutput(logFile)
	key, err := os.ReadFile("keyFile.data")
	if err != nil {
		log.Println(err)
	}
	ciphertext, err := os.ReadFile("myfile.data")
	if err != nil {
		log.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(plaintext))
}

