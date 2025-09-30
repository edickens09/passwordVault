package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	
	user "github.com/edickens09/passwordVault/user"
	connect "github.com/edickens09/passwordVault/connect"

)

func Menu() string {
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--------MENU--------")
	fmt.Println("1) Create new Entry")
	fmt.Println("2) Find specific Entry")
	fmt.Println("3) List all Entries")
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

// Maybe this should be rewritten so that it doesn't need the connection passed into it since most options don't need the connection. Maybe only establish a connection for the purpose of syncing server and client
func HandleCommands(conn net.Conn) {

	for {
	
		command := Menu()
		
		if command == "" {
			return
		}

		switch command {

		case "CREATE":
			HandleCreate()
			go connect.SyncToServer()

		case "STOP":
			fmt.Println("TCP client exit...")
			fmt.Fprintf(conn, command + "\n")
			return

		case "RETRIEVE":
			item := HandleRetrieve()
			if item == nil {
				fmt.Println("Unable to retrieve item due to error")
			}

			fmt.Println(item)
			continue

		case "LIST":
			HandleList()
			continue

		default:
			fmt.Println("Unknown Command: " + command)
			continue
		}
	}	
}

func main() {

	file, err := os.OpenFile("logs/clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
		return
	}

	log.SetOutput(file)
	
	//this needs sent to the server during the sync to get the correct information
	username, err := user.LoginUsername()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(username)

	c, err := connect.SyncFromServer()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	HandleCommands(c)
}
