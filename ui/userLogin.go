package ui

import (

	"github.com/edickens09/passwordVault/user"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m LoginText) Init() tea.Cmd {
	
	return textinput.Blink
}

func (m LoginText) View() tea.View {

	s := "Please enter Username and Password\n"

	for i := range m.inputs {
		s+=m.inputs[i].View() + "\n"
	}

	s += "\nPress q to quit"

	return tea.NewView(s)
}

func (m LoginText) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl + c", "q":
			return m, tea.Quit

		case "enter":

			userText := m.inputs[0]
			passText := m.inputs[1]

			userName := userText.Value()
			password := passText.Value()

			//I should have a check for username in database first then if user is in database check password
			//these are standins for logic
			user.CheckUserPath(userName)
			passHash, err := user.HashPassword(password, "123")
			if err != nil {
				//Need to put something here
				return m, nil
			}

			databaseHash = passHash + "123"
			if err = ComparePasswords(passHash, databaseHash); err != nil {
				return m, nil
			}

			mainMenu := MainMenu()
			return m, SwitchModel(mainMenu)

		case "tab", "shift+tab", "up", "down":
			if msg.String() == "up" || msg.String() == "tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			}

			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}
			
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
			m.inputs[i].Blur()
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.UpdateInputs(msg)

	return m, cmd
}

func LoginUser() LoginText {
	et := LoginText {
		inputs: make([]textinput.Model, 2),
	}

	var text textinput.Model

	text = textinput.New()
	text.Placeholder = "Username"
	text.Focus()
	text.CharLimit = 32
	text.SetWidth(20)
	et.inputs[0] = text

	text = textinput.New()
	text.Placeholder = "Password"
	text.CharLimit = 64
	text.SetWidth(20)
	et.inputs[1] = text

	return et
}

func (m *LoginText) UpdateInputs(msg tea.Msg) tea.Cmd {

	cmds := make([]tea.Cmd, len(m.inputs))
	
	//Only text inputs Focus() set will respond, so it's safe to simply
	//update all of them here without any further logic
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
