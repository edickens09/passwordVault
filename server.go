package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"errors"
	"log"
)


type Version struct {
	//Major version will break backwards compaibility
	Major uint8
	Minor uint16
	Patch uint16
}

//var protocolVersion = Version{Major:0, Minor:1, Patch:0}

//Handles anything for the initial Authentication
func HandleAuthentication(conn net.Conn) error {
	authenticationKey := "Authentication Key"

	//scanner, _ := bufio.NewReader(c).ReadString('\n')
	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() {
		return errors.New("Auth Error\n")
	}


	authenticationAttempt := scanner.Text()

	if authenticationAttempt != authenticationKey {
		return errors.New("Auth Error\n")
	}
	return nil
}

//Handles anything for the handshake
func HandleHandshake(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() {
		return errors.New("Connection Error\n")
	}

	handshakeLine := scanner.Text()
	if handshakeLine != "Version 0.001" {
		
		return errors.New("Version handshake Error\n")
	}

	return nil
}

func HandleConnection(c net.Conn) {

	defer c.Close()

	if err := HandleAuthentication(c); err != nil {
		log.Print(err)
		c.Write([]byte(err.Error()))
		return
	}

	c.Write([]byte("Success\n"))
	
	if err := HandleHandshake(c); err != nil {
		log.Println(err)
		c.Write([]byte(err.Error()))
		return
	}
	
	fmt.Println("handshake accepted")
	c.Write([]byte("handshake accepted\n"))
	
	// continuous handling of incoming packets
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
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

	file, err := os.OpenFile("serverLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error with log file")
	}

	log.SetOutput(file)

	PORT := ":19865"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go HandleConnection(c)
	}

}
