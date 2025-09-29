package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
	
	user "github.com/edickens09/passwordVault/user"
	connect "github.com/edickens09/passwordVault/connect"
	"gopkg.in/yaml.v3"

)

type Config struct {
	Host string `yaml:"server"`
	Port string `yaml:"port"`
}

type Version struct{
	//Major Verison number will break backwards compatibility
	Major uint8
	//Minor Version has new features or commands, server may support multiple version
	Minor uint16
	//Patch Version, should only have bug fixes and shouldn't break
	Patch uint16
}

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
func HandleCommands(conn net.Conn, username string) {

	for {
	
		command := Menu()
		
		if command == "" {
			return
		}

		switch command {

		case "CREATE":
			HandleCreate(username)
			go connect.SyncToServer()

		case "STOP":
			fmt.Println("TCP client exit...")
			fmt.Fprintf(conn, command + "\n")
			return

		case "RETRIEVE":
			item := HandleRetrieve(username)
			if item == nil {
				fmt.Println("Unable to retrieve item due to error")
			}

			fmt.Println(item)
			continue

		case "LIST":
			HandleList(username)
			continue

		default:
			fmt.Println("Unknown Command: " + command)
			continue
		}
	}	
}

func SyncFromServer() (net.Conn, error) {
	var config Config

	yFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error opening config file")
		return nil, err
	}

	err2 := yaml.Unmarshal(yFile, &config)
	if err2 != nil {
		fmt.Println("Error with config file")
		return nil, err2
	}

	server := config.Host
	port := config.Port

	c, err := net.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := connect.HandleAuthentication(c); err != nil {
		return nil, err
	}

	if err := connect.HandleHandshake(c); err != nil {
		return nil, err
	}

	return c, nil
}

func main() {

	file, err := os.OpenFile("logs/clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
		return
	}

	log.SetOutput(file)
	
	//this needs sent to the server during the sync to get the correct information
	username, err := user.GetUsername()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(username)

	c, err := SyncFromServer()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	HandleCommands(c, username)
}
