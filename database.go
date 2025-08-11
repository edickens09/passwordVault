package database

import(
	"fmt"
	"os"
	"errors"
	"bufio"
)

type Database struct {
	serviceName string
	username string
	password string
}

func RetrieveData(serviceName string, vault Database) error {
	file, err := os.OpenFile("vault.data",)
	if err != nil {
		return errors.New("Vault error\n")
	}
	
	data, err := ParseVault(serviceName, vault)
	if err != nil {
		return errors.New("Error with Vault\n")
	}

}

func ParseVault(name string, vault Database) ([]string, error) {

	file, err := os.OpenFile("vault.data")
	if err != nil {
		return nil, errors.New("Error with vault File")
	}
	scanner, err := bufio.NewScanner(file)
	if err != nil {
		return nil, errors.New("Vault error\n")
	}

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		if items[0] == name {
			return items, nil
		} else {
			return "Service Name not found", nil
		}
	}

}
