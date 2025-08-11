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

/*func RetrieveData(serviceName string, vault Database) error {
	file, err := os.OpenFile("vault.data",)
	if err != nil {
		return errors.New("Vault error\n")
	}
	
	data, err := ParseVault(serviceName, vault)
	if err != nil {
		return errors.New("Error with Vault\n")
	}

}*/

func ParseVault(name string, vault Database) ([]string, error) {

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
