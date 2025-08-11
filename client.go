package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"errors"
	"log"

)
type Version struct{
	//Major Verison number will break backwards compatibility
	Major uint8
	//Minor Version has new features or commands should mostly work
	Minor uint16
	//Patch Version, should only have bug fixes and shouldn't break
	Patch uint16
}

//var protocolVersion = Version{Major:0, Minor:1, Patch:0}

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
	version := "Version 0.001\n"
	fmt.Fprintf(conn, version)

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

	var data = Database{serviceName:"This", username:"is", password:"test"}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")

		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		switch command {

		case "STOP\n":
			fmt.Println("TCP client exit...")
			fmt.Fprintf(conn, command)
			return

		case "TEST\n":
			test := data.TestImport()
			fmt.Println(test)
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

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host.")
	}

	file, err := os.OpenFile("clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
	}

	log.SetOutput(file)

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT + ":19865")
	if err != nil {
		log.Println(err)
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
	/*for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")

		command, _ := reader.ReadString('\n')

		switch command {
		
			case "STOP":
			fmt.Println("TCP client exiting...")
			return

			default:
			fmt.Fprintf(c, command + "\n")
			message, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Error closing connection")
					return
				}
			}
			fmt.Println("->: " + message)
		}
	}*/
}

