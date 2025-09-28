package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"

	database "github.com/edickens09/passwordVault/database"
)

func HandleList() {

	err := database.ListVault()
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

	data, err := database.ParseVault(serviceName)
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

func HandleCreate(vault database.Database) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	if err := vault.CreateEntry(serviceName, "eric"); err != nil {
		fmt.Println("\nVault error check logs for more detail")
		log.Fatalln(err)
		return
	}
}
