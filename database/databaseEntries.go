package database

import (
	"fmt"

	//"github.com/mattn/go-sqlite3"
)

type Entry struct {

	Name string
	EntryType string
	Username string
	EncryptedPassword string
	PasswordKey string
	WebAddress string
	CreationDate string
	ModifiedDate string
	Trash bool
	Comments string
}

func CreateEntry(NewEntry Entry) error {

	/*
	NewEntry.Name enters database
	NewEntry.EntryType enters database
	NewEntry.Username enters database
	NewEntry.EncryptedPassword enters database
	NewEntry.PasswordKey enters database
	NewEntry.WebAddress enters database
	NewEntry.CreationData enters database
	NewEntry.ModifiedDate enters database
	NewEntry.Comments enters database
	*/


	return nil
}
