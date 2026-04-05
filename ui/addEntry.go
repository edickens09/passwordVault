package ui

import (

	"github.com/edickens09/passwordVault/connect"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func CreateEntry() tea.Msg {

	//database.HandleCreate
	go connect.SyncToServer()

	return Msg("")
}

func (m EntryText) Init() tea.Cmd {
	return textinput.Blink
}

func (m EntryText) View() tea.View {
	var s string
	s = "Entry menu Working"
	
	return tea.NewView(s)
}

func (m EntryText) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "tab", "shift+tab", "up", "down":
			if msg.String() == "up" || msg.String() == "up" {
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

		}
	}
	return m, nil
} 

func InitialModel() EntryText {

	et := EntryText {
		inputs: make([]textinput.Model, 6),
	}

	var text textinput.Model
	text = textinput.New()
	text.Placeholder = "Name"
	text.Focus()
	et.inputs[0] = text

	text = textinput.New()
	text.Placeholder = "Entry Type"
	et.inputs[1] = text

	text = textinput.New()
	text.Placeholder = "Username"
	et.inputs[2] = text

	text = textinput.New()
	text.Placeholder = "Password"
	et.inputs[3] = text

	text = textinput.New()
	text.Placeholder = "Web or ip address"
	et.inputs[4] = text

	text = textinput.New()
	text.Placeholder = "Comments"
	et.inputs[5] = text

	return et
}
