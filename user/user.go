package user

import (
	"fmt"
	"errors"
	"os"
	"log"
)

var Username = "" 

// this needs a way to verify to the server as well as locally. 
// if using same for both server side doesn't necessarily need user.CheckUserPath()

/* does a check if the database exists, if that doesn't return an error then parses through the database
if that doesn't return an error then it check the password in database in comparision to the password input
if the salted hashes match then it will return nil*/

func CheckUserPath(username string) {

	if _, err := os.Stat("user/" + username + "/"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("user directory not found creating now")
		if err := os.MkdirAll("user/" + username + "/", os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}
