package ui

import (
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
	WebAddress string
	Username string
	Password string
	CreationDate string
	ModifiedDate string
	Trash bool
	Comments string
}
