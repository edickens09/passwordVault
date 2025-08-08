package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"errors"
)

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
		if err == io.EOF {
			fmt.Println("Connection closed. Exiting")
			return err
		}
	}

	fmt.Println(handshakeAnswer)
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host.")
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT + ":19865")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := HandleAuthentication(c); err != nil {
		fmt.Println(err)
		return
	}

	if err := HandleHandshake(c); err != nil {
		fmt.Println(err)
		return
	}


	for {
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
	}
}

