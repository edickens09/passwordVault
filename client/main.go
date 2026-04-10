package main

import (
	"fmt"
	"os"
	"log"
	
//	user "github.com/edickens09/passwordVault/user"
	ui "github.com/edickens09/passwordVault/ui"

)

func main() {

	file, err := os.OpenFile("logs/clientLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Log File Error")
		return
	}

	log.SetOutput(file)
	
	ui.StartApp()
}
