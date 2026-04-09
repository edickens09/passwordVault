package ui

import (

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m UserText) Init() tea.Cmd {
	
	return textinput.Blink
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

		case "enter":
			return m, tea.Quit

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
	text.SetWidth(20)
	et.inputs[0] = text

	text = textinput.New()
	text.Placeholder = "Password"
	text.CharLimit = 64
	text.SetWidth(20)
	et.inputs[1] = text

	return et
}
