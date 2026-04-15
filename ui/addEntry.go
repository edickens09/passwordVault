//Package ui handles the ui menus and inputs
package ui

import (

	"time"

	//"github.com/edickens09/passwordVault/connect"
	"github.com/edickens09/passwordVault/database"
	encrypt "github.com/edickens09/passwordVault/encryption"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

/*func CreateEntry() tea.Msg {

	//database.HandleCreate
	go connect.SyncToServer()

	return Msg("")
}*/

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

		case "enter":
			
			var NewEntry database.Entry
                        var entryName string
                        var entryType string

			today := time.Now().UTC()

			if entryName = m.inputs[0].Value(); entryName == "" {
				//change ui to indicate that blank entry isn't allowed for name
			}

			if entryType = m.inputs[1].Value(); entryType == "" {
				//change ui to indicate that blank entry isn't allowed for type
				//if no type matches use "other" typed
			}

			entryUsername := m.inputs[2].Value()
			entryPassword := m.inputs[3].Value()
			entryURL := m.inputs[4].Value()
			entryComments := m.inputs[5].Value()

			encryptedPassword, key, err := encrypt.EncryptPassword(entryPassword)
			if err != nil {
				//I want to add this error to a log file in addition to changing the ui screen to indicate failure
				return m, nil 
			}

			NewEntry.Name = entryName
			NewEntry.EntryType = entryType
			NewEntry.Username = entryUsername
			NewEntry.EncryptedPassword = encryptedPassword
			NewEntry.WebAddress = entryURL
			NewEntry.Comments = entryComments
			NewEntry.CreationDate = today.Format(time.UnixDate)
			NewEntry.ModifiedDate = today.Format(time.UnixDate)

		if err := database.CreateEntry(NewEntry); err != nil {

				// I want to add this error to a log file in addition to changing to ui screen to indicate failure
				return m, nil
			}

			mainMenu := MainMenu()
			return m, SwitchModel(mainMenu)
		}
	}

	cmd := m.UpdateInputs(msg)

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
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[0] = text

	text = textinput.New()
	text.Placeholder = "Entry Type"
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[1] = text

	text = textinput.New()
	text.Placeholder = "Username"
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[2] = text

	text = textinput.New()
	text.Placeholder = "Password"
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[3] = text

	text = textinput.New()
	text.Placeholder = "Web or ip address"
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[4] = text

	text = textinput.New()
	text.Placeholder = "Comments"
	text.CharLimit = 64
	text.SetWidth(32)
	et.inputs[5] = text

	return et
}

func (m *EntryText) UpdateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
