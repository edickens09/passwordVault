package user

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"log"
	"strings"
)

var Username = "" 

// this needs a way to verify to the server as well as locally. 
// if using same for both server side doesn't necessarily need user.CheckUserPath()
func LoginUsername() error {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")

	userName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	Username = strings.TrimSuffix(userName, "\n")

	if Username == "" {
		fmt.Println("Username cannot be an empty string")
		if err := LoginUsername(); err != nil {
			return err
		}
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

func GetUsername() (string, error) {
	return "eric", nil

}
