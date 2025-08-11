package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"bufio"
	"log"
)

func main() {

	logFile, err := os.ReadFile("encryptionLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error with log file")
	}

	log.SetOutput(logFile)

	fmt.Println("Encryption Program v0.0001")

	fmt.Println("input text")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	//make the byte information for both the text input as well as the random generated key
	text := []byte(line)
	key := make([]byte, 32)
	_, err = rand.Read(key)
	err = os.WriteFile("keyFile.data", key, 0777)

	//generate cypther with 32 byte key
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
	}

	err = os.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Data Encrypted")

}
