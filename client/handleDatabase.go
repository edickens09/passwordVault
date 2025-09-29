package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"

	database "github.com/edickens09/passwordVault/database"
)

func HandleList(username string) {

	err := database.ListVault(username)
	if err != nil {
		fmt.Println("\nVault error check logs for more details")
		log.Println(err)
	}
	
}	

//this is not functional as of the new vault structure needs refactored in database package to make work again
func HandleRetrieve(username string) [] string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	data, err := database.ParseVault(serviceName, username)
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

func HandleCreate(username string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	if err := database.CreateEntry(serviceName, username); err != nil {
		fmt.Println("\nVault error check logs for more detail")
		log.Fatalln(err)
		return
	}
}
