package database

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
)
//keep a list of all the service names in the database and a location of the file with necessary info. This will make parsing and looking for services easier 
func initalizeDatabase(){
	/* should there be one database that has both the username and the encrypted password stored for comparision
	as well as the services that are with that username
	other option could be two seperate databases, one that is username with encrypted password and the other with 
	username and services that are attached to that username
*/
}
func HandleList() {

	err := ListVault()
	if err != nil {
		fmt.Println("\nVault error check logs for more details")
		log.Println(err)
	}
	
}

//this is not functional as of the new vault structure needs refactored in database package to make work again
func HandleRetrieve() [] string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	data, err := ParseVault(serviceName)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		return nil
	}

	if data == nil {
		fmt.Println("Vault entry not found")
	}

	return data
}

// this needs to be refactored so that it gets all the information and then passes it to a seperate function that creates the vault file not how it works currently.
func HandleCreate() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	fmt.Printf("What is the username for %v:\n", serviceName)

	username, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	username = strings.TrimSuffix(username, "\n")

	fmt.Printf("What is the password for %v:\n", serviceName)

	password, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	password = strings.TrimSuffix(password, "\n")


	if err := CreateEntry(serviceName, username, password); err != nil {
		fmt.Println("\nVault error check logs for more detail")
		log.Fatalln(err)
		return
	}
}
