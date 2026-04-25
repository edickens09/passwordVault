package user

import (
	"fmt"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string
	PasswordHash string
}

type Db struct {
	db *sql.DB
}

const UserDB = "userDb"

//parse through database to see if a usernamme exists. If it doesn't should return an error
func UserExists(username string, db *Db) (bool, error) {

	query := fmt.Sprintf("SELECT * FROM %s WHERE user = ?", UserDB)

	user, err := db.db.Query(query, username)
	if err != nil {
		return false, err
	}

	defer user.Close()

	return true, nil	
}

/* maybe this should just check to make sure the database exists and everything necessary
is in place instead of trying to edit anything
*/
func InitalizeDatabase(dBName string) (Db, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Db{}, err
	}

	dBAddress := filepath.Join(homeDir, fmt.Sprintf(".%s", dBName))

	db, err := sql.Open("sqlite3", dBAddress)
	if err != nil {
		return Db{}, err
	}

	sqlStr := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
			id TEXT NOT NULL PRIMARY KEY,
			user TEXT NOT NULL,
			passwordHash BYTE NOT NULL,
			
		);`, UserDB,
		)

	_, err = db.Exec(sqlStr)
	if err != nil {
		return Db{}, err
	}

	return Db{db}, nil

}

// add the username to the usertable, store it with the salt and hashed password
func AddUser(username string, db *Db, passwordHash string) error {

	sqlStr := fmt.Sprintf(`INSERT INTO %s (id, user, passwordHash)
		VALUES(?, ?, ?)`, UserDB)
	
	return nil
	
}

