package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
)

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

		text, _ := reader.ReadString('\n')

		switch text {
		
		case "STOP":
			fmt.Println("TCP client exiting...")
			return

		default:
			fmt.Fprintf(c, text + "\n")
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

