package database

import(
	"fmt"
	"os"
	"errors"
	"bufio"
	"strings"
)

type Database struct {
	serviceName string
	username string
	password string
}

func (data Database) ParseVault(name string) ([]string, error) {

	file, err := os.Open("vault.data")
	if err != nil {
		return nil, errors.New("error with vault File")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, errors.New("error reading vault")
	}

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		if items[0] == name {
			return items, nil
		}	
	}

	return nil, nil

}

func (data Database) ListVault() (error) {
	file, err := os.Open("vault.data")
	if err != nil {
		return errors.New("error with vault file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return errors.New("error reading vault")
	}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return nil
}

func (data Database) CreateEntry(name string) error {

	fmt.Printf("What's the username for %v\n", name)
	reader := bufio.NewReader(os.Stdin)

	username, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("error getting username")
	}

	username = strings.TrimSuffix(username, "\n")

	fmt.Printf("What's the password for %v\n", name)
	reader = bufio.NewReader(os.Stdin)

	password, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("error getting password")
	}

	password = strings.TrimSuffix(password, "\n")
	
	data.serviceName = name
	data.username = username
	data.password = password
	
	file, err := os.OpenFile("vault.data", os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		return errors.New("error opening file for writing")
	}
	
	defer file.Close()
	
	fmt.Fprintln(file, data)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
