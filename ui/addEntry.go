package ui

import (

	"github.com/edickens09/passwordVault/connect"
	tea "charm.land/bubbletea/v2"
)

func CreateEntry() tea.Msg {

	//database.HandleCreate
	go connect.SyncToServer()

	return Msg("")
}


