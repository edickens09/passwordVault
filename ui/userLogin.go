package ui

import (

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m UserText) Init() tea.Cmd {
	
	return nil 
}

func (m UserText) View() tea.View {
	s := "This view is working currently"

	return tea.NewView(s)
}

func (m UserText) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl + c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func LoginUser() UserText {
	et := UserText {
		inputs: make([]textinput.Model, 2),
	}

	var text textinput.Model

	text = textinput.New()
	text.Placeholder = "Username"
	text.Focus()
	text.CharLimit = 32
	et.inputs[0] = text

	text = textinput.New()
	text.Placeholder = "Password"
	text.CharLimit = 64
	et.inputs[1] = text

	return et
}
