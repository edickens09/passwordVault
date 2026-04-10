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

type EntryText struct {

	focusIndex int
	inputs []textinput.Model
}

type LoginText struct {

	focusIndex int
	inputs []textinput.Model
}

