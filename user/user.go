package user

import (
	"fmt"
	"errors"
	"os"
	"log"
)

var Username = "" 

// this needs a way to verify to the server as well as locally. 

/* does a check if the database exists, if that doesn't return an error then parses through the database
if that doesn't return an error then it check the password in database in comparision to the password input
if the salted hashes match then it will return nil*/

// CheckUserPath checks directory for user and creates it if not found
func CheckUserPath(username string) {

	if _, err := os.Stat("user/" + username + "/"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("user directory not found creating now")
		if err := os.MkdirAll("user/" + username + "/", os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}

func LoginUser(username string, password string) error {

	passHash, err := HashPassword(password, []byte("123"), []byte("123"))
	if err != nil {
		return err
	}

	if err := ComparePasswords(username, passHash); err != nil {
		return err
	}
	//takes the username and passwordHash if they match the entry in the database then return nil

	return nil
}
