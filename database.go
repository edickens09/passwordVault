package main

import(
//	"fmt"
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
		return nil, errors.New("Error with vault File")
	}

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, errors.New("Error reading vault")
	}

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		if items[0] == name {
			return items, nil
		}	
	}

	return nil, nil

}

func (data Database) TestImport() string{
	string := "This is only here to test the import is working"
	return string
}
