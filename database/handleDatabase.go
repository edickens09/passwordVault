package database

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
)

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

func HandleCreate() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	if err := CreateEntry(serviceName); err != nil {
		fmt.Println("\nVault error check logs for more detail")
		log.Fatalln(err)
		return
	}
}
