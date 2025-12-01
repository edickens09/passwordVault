package user

import (
	"fmt"
	"database/sqlite"
	"os"
	"errors"
)

//parse through database to see if a usernamme exists. If it doesn't should return an error
func ParseDatabase() error {

	fmt.Println("This is working to theis point")

	return nil
}

func DatabaseUserExists() error {

	fmt.Println("This is working to this point")

	return nil	
}

func InitalizeDatabase(userDatabase string) error {

	if _, err := os.Stat(userDatabase); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(userDatabase, os.ModePerm); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil

}
