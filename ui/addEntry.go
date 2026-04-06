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
	s = "Entry menu Working\n\n"

	for i := range m.inputs {
		s += m.inputs[i].View() + "\n"
	}

	s += "\nPress q to quit"
	
	return tea.NewView(s)
}

func (m EntryText) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "tab", "shift+tab", "up", "down":
			if msg.String() == "up" || msg.String() == "shift+tab" {
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
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
					}
				// Remove focused state
				m.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)

		case "q", "ctrl + c":
			return m, tea.Quit

		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
} 

func InitialModel() EntryText {

	et := EntryText {
		inputs: make([]textinput.Model, 6),
	}

	var text textinput.Model
	text = textinput.New()
	text.Placeholder = "Name"
	text.Focus()
	text.CharLimit = 156
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

func (m *EntryText) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
