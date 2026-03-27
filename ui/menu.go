package ui

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func MainMenu() RootMenu {
	return RootMenu {
		choices: []string{"Entries", "Check Connection", "Settings", "Logout", "Exit"},

	}
}

func (m RootMenu) Init() tea.Cmd {

	return nil
}

func (m RootMenu) View() tea.View {
	s := "Please make a selection:\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = "*"
		}

		s += fmt.Sprintf("[%s] %s\n", cursor,choice)
	}

	s += "\nPress q to quit\n"

	return tea.NewView(s)
}

func (m RootMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit 

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case "enter":
			choice := m.choices[m.cursor]

			switch choice {
			case "Entries":
				return m, createEntry

			case "Check Connection":
				return m, nil

			case "Settings":
				return m, nil

			case "Logout":
				return m, nil
			
			case "Exit": 
				return m, tea.Quit
			}
		}


	}
	return m, nil
}

func StartApp() {
	p := tea.NewProgram(MainMenu())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error %v", err)
		os.Exit(1)
	}
}


