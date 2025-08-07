package main

import (
	"bufio"
	"fmt"
	"net"
//	"os"
	"strings"
	"errors"
)

func HandleAuthentication(c net.Conn) error {
	fmt.Println("Made it to the Authentication step successfully")
	authenticationKey := "Authentication Key"

	//scanner, _ := bufio.NewReader(c).ReadString('\n')
	scanner := bufio.NewScanner(c)
	fmt.Println(scanner)

/*	if !scanner.Scan() {
		return errors.New("Auth Error\n")
	}*/


	authenticationAttempt := scanner.Text()
	fmt.Println(authenticationAttempt)


	if authenticationAttempt != authenticationKey {
		return errors.New("Auth Error\n")
	}
	return nil
}

func HandleConnection(c net.Conn) {

	defer c.Close()

	if err := HandleAuthentication(c); err != nil {
		fmt.Print(err)
		c.Write([]byte(err.Error()))
		return
	}

	c.Write([]byte("Success\n"))
	
	// starting communication for handshake
	scanner := bufio.NewScanner(c)
	fmt.Println(scanner)

	//testing if scanner does exist should probably expand upon this later
	if !scanner.Scan() {
		return
	}
	//version handshake 	
	handshakeLine := scanner.Text()
	fmt.Println(handshakeLine)
	if handshakeLine != "Version 0.001" {
		c.Write([]byte("Incompatible client. Closing connection\n"))
		fmt.Println("Incompatible client. Closing connection")
		return
	}
	
	fmt.Println("handshake accepted")
	c.Write([]byte("handshake accepted\n"))
	
	// continuous handling of incoming packets
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))

		switch temp {

		case "CREATE":
			fmt.Println("Creating New Entry")
			HandleCreate(c, "PasswordFile.data")
			continue

		case "RETRIEVE":
			fmt.Println("Sending Entry")
			HandleRetrieve(c)
			continue

		case "LIST":
			fmt.Println("Sending List")
			HandleList(c)
			continue

		case "STOP":
			fmt.Println("Stopping a connection")
			return

		default:
			fmt.Println(temp)
			c.Write([]byte("Unknown Command\n"))
		}
	}
}

func HandleRetrieve(c net.Conn) {
	// replace with retriving specific account information
	c.Write([]byte("Retrieve is working to this point\n"))
}

func HandleCreate(c net.Conn, file string){
	// need to replace this with one that creates an entry in a file that starts with the service name, then username, then encrypted password
	c.Write([]byte("Create is working to this point\n"))
}

func HandleList(c net.Conn) {
	// replace with return a list with all account names not user names the name of the service
	c.Write([]byte("List is working to the point\n"))
}

func main() {

	PORT := ":19865"
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
