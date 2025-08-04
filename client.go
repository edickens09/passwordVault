package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"io"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, "Version 0.001\n")
	/*
	userInput := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := userInput.ReadString('\n')

	fmt.Println(username)
	fmt.Fprintf(c, username + "\n")
	*/

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
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}

