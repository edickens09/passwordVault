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
	fmt.Println("Working here")
	fmt.Fprintf(c, authenticationKey)
	fmt.Println(authenticationKey)
	fmt.Println("Working here too")

	authenticationAttempt, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("Working here 3")
	
	if authenticationAttempt != "Success" {
		return errors.New(authenticationAttempt)
	}else{
		return nil
	}
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

	fmt.Println("Authentication Successful")

	fmt.Fprintf(c, "Version 0.001\n")

	handshakeAnswer, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("Connection closed. Exiting")
			return
		}
	}
	fmt.Println(handshakeAnswer)


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

