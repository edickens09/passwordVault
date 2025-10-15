package main

import (
	"fmt"
	"os"
	"log"
	
	user "github.com/edickens09/passwordVault/user"
	connect "github.com/edickens09/passwordVault/connect"
	ui "github.com/edickens09/passwordVault/ui"

)

func main() {

	file, err := os.OpenFile("logs/clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
		return
	}

	log.SetOutput(file)
	
	//this needs sent to the server during the sync to get the correct information
	if err := user.LoginUser(); err != nil {
		log.Fatalln(err)
	}

	c, err := connect.SyncFromServer()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	ui.HandleCommands(c)
}
