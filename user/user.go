package user

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"log"
	"strings"
	"database/sqlite"
)

var Username = "" 

// this needs a way to verify to the server as well as locally. 
// if using same for both server side doesn't necessarily need user.CheckUserPath()

/* does a check if the database exists, if that doesn't return an error then parses through the database
if that doesn't return an error then it check the password in database in comparision to the password input
if the salted hashes match then it will return nil*/
func LoginUser() error {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")

	userName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	Username = strings.TrimSuffix(userName, "\n")
	
	//this is a recursive statement that prevents the user string from being empty
	if Username == "" {
		fmt.Println("Username cannot be an empty string")
		if err := LoginUser(); err != nil {
			return err
		}
	}

	fmt.Print("Password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	password = strings.TrimSuffix(password, "\n")

	passwordHash, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error Hashing password")
	}

	if err := ComparePasswords(passwordHash); err != nil {
		fmt.Println("Invalid Username or Password")
	}

	CheckUserPath(Username)

	return nil

}

func CheckUserPath(username string) {

	if _, err := os.Stat("user/" + username + "/"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("user directory not found creating now")
		if err := os.MkdirAll("user/" + username + "/", os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}
