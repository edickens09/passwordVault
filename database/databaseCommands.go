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

func CreateEntry(name string) error {

	var data Database

	fmt.Printf("What's the username for %v\n", name)
	reader := bufio.NewReader(os.Stdin)

	serviceUsername, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("error getting username")
	}

	serviceUsername = strings.TrimSuffix(serviceUsername, "\n")

	fmt.Printf("What's the password for %v\n", name)
	reader = bufio.NewReader(os.Stdin)

	password, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("error getting password")
	}

	password = strings.TrimSuffix(password, "\n")

	passwordHash, err := EncryptPassword(password)
	if err != nil {
		return errors.New("error hashing password")
	}
	
	data.username = serviceUsername
	data.password = passwordHash

	username := user.Username
	
	file, err := os.OpenFile("user/" + username + "/" + name + ".vault", os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error opening file for writing")
	}
	
	defer file.Close()
	
	fmt.Fprintln(file, data)
	//find a way to check and make sur eit wrote and created an error

	return nil
}
