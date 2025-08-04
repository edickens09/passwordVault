package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Encryption Program v0.0001")

	fmt.Println("input text")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	//make the byte information for both the text input as well as the random generated key
	text := []byte(line)
	key := make([]byte, 32)
	_, err = rand.Read(key)
	err = os.WriteFile("keyFile.data", key, 0777)

	//generate cypther with 32 byte key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data Encrypted")

}
