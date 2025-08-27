package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"errors"
	"log"
	"strings"
	"encoding/binary"

	"github.com/edickens09/passwordVault/database"
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

	fmt.Println("--------MENU--------")
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

		case "4\n":
			return "STOP"

		default:
			fmt.Println("That option doesn't exist")
			continue
		}

	}

}

func HandleAuthentication(c net.Conn) error {

	authenticationKey := "Authentication Key\n"
	fmt.Fprintf(c, authenticationKey)

	authenticationAttempt, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return err
	}

	if authenticationAttempt != "Success\n" {
		return errors.New(authenticationAttempt)
	}else{
		fmt.Println("Authentication Successful")
		return nil
	}
}

func HandleHandshake(conn net.Conn) error {

	clientVer := Version {
		Major:00,
		Minor:01,
		Patch:01,
	}

	err := binary.Write(conn, binary.BigEndian, clientVer)

	handshakeAnswer, err := bufio.NewReader(conn). ReadString('\n')
	if err != nil {
		log.Println(err)
		if err == io.EOF {
			fmt.Println("Connection closed. Exiting")
			return err
		}
	}

	fmt.Println(handshakeAnswer)
	return nil
}

func HandleCommands(conn net.Conn) {

	//This feels wrong, I need to initalize the database, but I don't think this is the correct way to do it
	//Should I look at a different way to do this?
	var data = database.Database{}
	for {
		/*reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")

		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}*/

		command := Menu()
		
		if command == "" {
			return
		}

		switch command {

		case "CREATE":
			HandleCreate(data)

		case "STOP":
			fmt.Println("TCP client exit...")
			fmt.Fprintf(conn, command)
			return

		case "RETRIEVE":
			item := HandleRetrieve(data)

			fmt.Println(item)
			continue


		default:
			fmt.Fprintf(conn, command)
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err)
				if err == io.EOF {
					fmt.Println("Error closing connection")
					return
				}
			}
			
			fmt.Println("-> " + message)
		}
	}	
}

func HandleRetrieve(vault database.Database) [] string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	data, err := vault.ParseVault(serviceName)
	if err != nil {
		log.Println(err)
	}

	if data == nil {
		fmt.Println("Vault Entry not found")
	}

	return data
}

func HandleCreate(vault database.Database) {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Service Name: ")

	serviceName, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	serviceName = strings.TrimSuffix(serviceName, "\n")

	if err := vault.CreateEntry(serviceName); err != nil {
		log.Println(err)
	}
}

func main() {

	var config Config

	yFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error opening config file")
	}

	err2 := yaml.Unmarshal(yFile, &config)
	if err2 != nil {
		fmt.Println("Error with config file")
	}

	server := config.Host
	port := config.Port

	file, err := os.OpenFile("logs/clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
	}

	log.SetOutput(file)

	c, err := net.Dial("tcp", server + ":" + port)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return
	}

	if err := HandleAuthentication(c); err != nil {
		log.Println(err)
		return
	}

	if err := HandleHandshake(c); err != nil {
		log.Println(err)
		return
	}

	HandleCommands(c)
}

