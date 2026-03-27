package ui

type RootMenu struct {
	choices [] string
	cursor int
	selected map[int]struct{}
}

type EntryMenu struct {
	choices [] string
	cursor int
	selected map[int]struct{}
}

type Msg string
