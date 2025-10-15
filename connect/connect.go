package connect

import (
	"fmt"
	"log"
	"net"
	"io"
	"encoding/binary"
	"errors"
	"bufio"
	"os"

	"gopkg.in/yaml.v3"
)

type Version struct {
	//Major version number very likely to break backwards compatibility
	Major uint8
	// Minor version number may have new features, but server may still support old version
	Minor uint16
	// Patch version number should only include bug fixes and should break compatibility
	Patch uint16
}

type Config struct {
	Server string `yaml:"server"`
	Port string `yaml:"port"`
}

func HandleHandshake(conn net.Conn) error {
	clientVer := Version {
		Major:00,
		Minor:01,
		Patch:01,
	}

	err := binary.Write(conn, binary.BigEndian, clientVer)
	if err != nil {
		log.Println(err)
		return err
	}

// look closer into this. I think it could be cleaned up with the error return this works but I think it's wrong
	handshakeAnswer, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err)
		if err == io.EOF {
			fmt.Println("Connection closed. Exiting")
			return err
		}
	}
	log.Println(handshakeAnswer)

	return nil
}

func HandleAuthentication(conn net.Conn) error {

	authenticationKey := "Authentication Key\n"
	fmt.Fprintf(conn, authenticationKey)

	authenticationAttempt, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	if authenticationAttempt != "Success\n" {
		return errors.New(authenticationAttempt)
	}else {
		return nil
	}
}

func SyncToServer() {
	fmt.Println("simulating syncing to server is working")
}

func SyncFromServer() (net.Conn, error) {
	var config Config

	yFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error opening config file")
		return nil, err
	}

	err2 := yaml.Unmarshal(yFile, &config)
	if err2 != nil {
		fmt.Println("Error with config file")
		return nil, err2
	}

	server := config.Server
	port := config.Port

	c, err := net.Dial("tcp", server + ":" + port)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := HandleAuthentication(c); err != nil {
		return nil, err
	}

	if err := HandleHandshake(c); err != nil {
		return nil, err
	}

	return c, nil
}
