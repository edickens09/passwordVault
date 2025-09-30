package ui

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"log"

	connect "github.com/edickens09/passwordVault/connect"
	database "github.com/edickens09/passwordVault/database"
)

func HandleCommands(conn net.Conn) {

	for {

		command := Menu()

		if command == "" {
			return
		}

		switch command {

		case "CREATE":
			database.HandleCreate()
			go connect.SyncToServer()
			continue

		case "RETRIEVE":
			item := database.HandleRetrieve()
			if item == nil {
				fmt.Println("Unable to retrieve item due to error")
			}
			continue

		case "LIST":
			database.HandleList()
			continue

		case "STOP":
			fmt.Println("TCP Client exit...")
			fmt.Fprintln(conn, command + "\n")
			return

		default:
			fmt.Println("Unknown Command: " + command)
			continue
		}
	}
}

func Menu() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--------MENU--------")
	fmt.Println("1) Create New Entry ")
	fmt.Println("2) Find specific entry")
	fmt.Println("3) List all entries")
	fmt.Println("4) Exit")

	for {
		fmt.Print(">> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return ""
		}

		switch command {

		case "1\n":
			return "CREATE"

		case "2\n":
			return "RETRIEVE"

		case "3\n":
			return "LIST"

		case "4\n", ":q\n":
			return "STOP"

		default:
			fmt.Println("That option doesn't exist")
			continue
		}
	}
}
