package database

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"

	//"github.com/mattn/go-sqlite3"
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


