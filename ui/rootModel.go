package ui

import (

//	"fmt"

	"charm.land/bubbletea/v2"

)

func (r RootModel) Init() tea.Cmd {

	return r.activeModel.Init()

}

func (r RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(switchMsg); ok {
		r.activeModel = msg.next
		return r, r.activeModel.Init()
	}

	var cmd tea.Cmd
	r.activeModel, cmd = r.activeModel.Update(msg)
	
	return r, cmd

}

func (r RootModel) View() tea.View {

	return r.activeModel.View()
}

func RootStart() RootModel {

	return RootModel {
		activeModel: MainMenu(),
	}

}

func SwitchModel(m tea.Model) tea.Cmd {
	return func() tea.Msg{
		return switchMsg{next: m}
	}
}
