package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func HandleConnection(c net.Conn) {

	defer c.Close()
	
	// starting communication for handshake
	scanner := bufio.NewScanner(c)

	//testing if scanner does exist should probably expand upon this later
	if !scanner.Scan() {
		return
	}
	//version handshake 	
	handshakeLine := scanner.Text()
	if handshakeLine != "Version 0.001" {
		c.Write([]byte(string("Incompatible client. Closing connection")))
		fmt.Println("Incompatible client. Closing connection")
		return
	}
	
	fmt.Println("handshake accepted")
	c.Write([]byte(string("handshake accepted\n")))
	
	// continuous handling of incoming packets
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))

		switch temp {

		case "TEST":
			fmt.Println("Reading from a file")
			secretInfo, err := os.ReadFile("Secret Info.txt")
			if err != nil {
				fmt.Println(err)
			}

			c.Write([]byte(string(secretInfo)))
			continue

		case "STOP":
			fmt.Println("Stopping a connection")
			return

		default:
			fmt.Println(temp)
			c.Write([]byte(string("Unknown Command\n")))
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
	}
	
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go HandleConnection(c)
	}

}
