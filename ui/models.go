package ui

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
)

type BaseMenu struct {
	choices [] string
	cursor int
	selected map[int]struct{}
}

type EntryMenu struct {
	choices [] string
	cursor int
	selected map[int]struct{}
}

type RootModel struct {
	activeModel tea.Model
}

type Msg string

type switchMsg struct {
	next tea.Model
}

type Entry struct {

	Name string
	EntryType string
	Username string
	Password string
	WebAddress string
	//These need to be apart of the database, but not a part of the Entry struct maybe??
	//CreationDate string
	//ModifiedDate string
	//Trash bool
	Comments string
}

type EntryText struct {

	focusIndex int
	inputs []textinput.Model
}
