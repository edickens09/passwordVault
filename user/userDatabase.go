package user

import (
	"fmt"
	"database/sql"
	"os"
	"errors"
	"github.com/mattn/go-sqlite3"
)

//parse through database to see if a usernamme exists. If it doesn't should return an error
func PasswordParse(username string, databaseLocation string) error {
	db, err := sql.Open("sqlite3", databaseLocation)
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("This is working to theis point")

	return nil
}

func UserExists(username string, databaseLocation string) error {
	db, err := sql.Open("sqlite3", databaseLocation)
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("This is working to this point")

	return nil	
}

/* maybe this should just check to make sure the database exists and everything necessary
is in place instead of trying to edit anything
*/
func InitalizeDatabase(databaseLocation string) error {

	if _, err := os.Stat(databaseLocation); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(databaseLocation, os.ModePerm); err != nil {
			return err
		}
	} else {
		return err
	}

	db, err := sql.Open("sqlite3", databaseLocation)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil

}

// add the username to the usertable, store it with the salt and hashed password
func AddUser(username string, databaseLocation string) error {
	db, err := sql.Open("sqlite3", databaseLocation)
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

