package user

import (
	"fmt"
	"bufio"
	//"errors"
	"os"
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

	if userName == "\n" {
		fmt.Println("Username cannot be an empty string")
		userName, err = GetUsername()
		if err != nil {
			return "", err
		}
	}

	//user.CheckUserPath(userName)

	return userName, nil

}

func CheckUserPath(username string) {
	/* if path username != true {
		create path
	}*/
}
