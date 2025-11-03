package database

import(
	"fmt"
	"os"
	"errors"
	"bufio"
	"strings"

	user "github.com/edickens09/passwordVault/user"
)

type Database struct {
	username string
	password string
	key []byte
}

// needs refactored no longer using "vault.data"
func ParseVault(name string) ([]string, error) {

	file, err := os.Open("vault.data")
	if err != nil {
		return nil, errors.New("error with vault File")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		if items[0] == name {
			return items, nil
		}	
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("error reading vault")
	}

	return nil, nil

}

//needs refactoring not longer uses "vault.data"
func ListVault() (error) {
	file, err := os.Open("vault.data")
	if err != nil {
		return errors.New("error with vault file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		return errors.New("error reading vault")
	}

	return nil
}

func CreateEntry(name string, serviceUsername string, password string) error {

	var data Database

	passwordHash, key, err := EncryptPassword(password)
	if err != nil {
		return err
	}
	
	data.username = serviceUsername
	data.password = passwordHash
	data.key = key

	//putting this here for future reference this is the username for password manager itself. not for the entry being created
	username := user.Username
	
	file, err := os.OpenFile("user/" + username + "/" + name + ".vault", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error opening file for writing")
	}
	
	defer file.Close()
	
	fmt.Fprintln(file, data.username, data.password)
	keyString := string(key)
	fmt.Fprintln(file, keyString)
	//find a way to check and make sur eit wrote and created an error

	return nil
}
