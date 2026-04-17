package database

import(
)

type Database struct {
	username string
	password string
	key []byte
}

// needs refactored no longer using "vault.data"
func ParseVault(name string) ([]string, error) {
	
	return nil, nil

}

//needs refactoring not longer uses "vault.data"
func ListVault() (error) {

	return nil
}

func InitalizeDatabase() error {

	return nil
}
