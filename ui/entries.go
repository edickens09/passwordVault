package ui

import (
	//"github.com/edickens09/passwordVault/database"
	"github.com/edickens09/passwordVault/connect"

	tea "charm.land/bubbletea/v2"
)

func (m EntryMenu) Init () tea.Cmd {
	return nil
}

func (m EntryMenu) View () tea.View {

	s := "This seems to be working"
	return tea.NewView(s)
}
func (m EntryMenu) Update (msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl + c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor -- 
			}
		case "down", "j":
			if m.cursor < len(m.choices) -1 {
				m.cursor ++
			}

		case "Enter":
			choice := m.choices[m.cursor]

			switch choice {

			case "Add Entry":
				return m, nil

			case "Search":
				return m, nil

			case "List":
				return m, nil
			}
		}
	}

	return m, nil

}
func createEntry() tea.Msg {

	//database.HandleCreate()
	go connect.SyncToServer()

	return Msg("")
}


