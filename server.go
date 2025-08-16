package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"errors"
	"log"
	"encoding/binary"
)


type Version struct {
	//Major version will break backwards compaibility
	Major uint8
	Minor uint16
	Patch uint16
}

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

	serverVer := Version{
		Major:00,
		Minor:01,
		Patch:01,
	}
	var clientVer Version

	err := binary.Read(conn, binary.BigEndian, &clientVer )
	if err != nil {
		return errors.New("Version Connection Error\n")
	}

	if serverVer != clientVer {
		return errors.New("Version Compatibility Error\n")
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
			c.Write([]byte("Create Received\n"))
			continue

		case "RETRIEVE":
			fmt.Println("Sending Entry")
			c.Write([]byte("Retrieve Received\n"))
			continue

		case "LIST":
			fmt.Println("Sending List")
			c.Write([]byte("Here is list working\n"))
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

func main() {

	logFile, err := os.OpenFile("logs/serverLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error with log file")
	}

	log.SetOutput(logFile)

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
