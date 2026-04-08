package ui

import (

	"fmt"

	//"github.com/edickens09/passwordVault/database"
	//"github.com/edickens09/passwordVault/connect"

	tea "charm.land/bubbletea/v2"
)

func (m EntryMenu) Init() tea.Cmd {
	return nil
}

func (m EntryMenu) View() tea.View {

	s := "Please make a selection:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		
		if m.cursor == i {
			cursor = "*"
		}

		s += fmt.Sprintf("[%s] %s\n", cursor,choice)
	}

	s += "\nPress q to quit"

	return tea.NewView(s)
}

func (m EntryMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor -- 
			}
		case "down", "j":
			if m.cursor < len(m.choices) -1 {
				m.cursor ++
			}

		case "enter":
			choice := m.choices[m.cursor]

			switch choice {

			case "Add Entry":
				initialModel := InitialModel()
				return m, SwitchModel(initialModel)

			case "Search":
				return m, nil

			case "List":
				return m, nil

			case "Back":
				mainMenu := MainMenu()
				return m, SwitchModel(mainMenu)
			}
		}
	}

	return m, nil

}

func EntriesMenu() EntryMenu {

	return EntryMenu {
		choices: []string{"Add Entry", "Search", "List", "Back"},
	}
}

