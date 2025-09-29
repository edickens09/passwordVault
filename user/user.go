package user

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"log"
	"strings"
)

// this needs a way to verify to the server as well as locally. 
// if using same for both server side doesn't necessarily need user.CheckUserPath()
func GetUsername() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")

	userName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	userName = strings.TrimSuffix(userName, "\n")

	if userName == "" {
		fmt.Println("Username cannot be an empty string")
		userName, err = GetUsername()
		if err != nil {
			return "", err
		}
	}

	CheckUserPath(userName)

	return userName, nil

}

func CheckUserPath(username string) {

	if _, err := os.Stat("user/" + username + "/"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("user directory not found creating now")
		if err := os.MkdirAll("user/" + username + "/", os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
	/* if path username != true {
		create path
	}*/
}
