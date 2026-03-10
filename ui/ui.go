package ui

import (
	"fmt"
	"bufio"
	"net"
	"os"
	"log"

	connect "github.com/edickens09/passwordVault/connect"
	database "github.com/edickens09/passwordVault/database"

	tea "charm.land/bubbletea/v2"
)

type menu struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func InitialMenu() menu {
	return menu {
		choices: []string{"Create New Entry", "Search", "Display All",},

	}
}

func (m menu) Init() tea.Cmd {

	return nil
}

func (m menu) View() tea.View {
	s := "Please make a selection:\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = "*"
		}

		s += fmt.Sprintf("[%s] %s\n", cursor,choice)
	}

	s += "\nPress q to quit\n"

	return tea.NewView(s)
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit 

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		}


	}
	return m, nil
}

func StartApp() {
	p := tea.NewProgram(InitialMenu())
	if _, err := p.Run(); err != nil {
		fmt.Println("There has been an error %v", err)
		os.Exit(1)
	}
} 

func HandleCommands(conn net.Conn) {

	for {

		command := Menu()

		if command == "" {
			return
		}

		switch command {

		case "CREATE":
			database.HandleCreate()
			go connect.SyncToServer()
			continue

		case "RETRIEVE":
			item := database.HandleRetrieve()
			if item == nil {
				fmt.Println("Unable to retrieve item due to error")
			}
			continue

		case "LIST":
			database.HandleList()
			continue

		case "STOP":
			fmt.Println("TCP Client exit...")
			fmt.Fprintln(conn, command + "\n")
			return

		default:
			fmt.Println("Unknown Command: " + command)
			continue
		}
	}
}

func Menu() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--------MENU--------")
	fmt.Println("1) Create New Entry ")
	fmt.Println("2) Find specific entry")
	fmt.Println("3) List all entries")
	fmt.Println("4) Exit")

	for {
		fmt.Print(">> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return ""
		}

		switch command {

		case "1\n":
			return "CREATE"

		case "2\n":
			return "RETRIEVE"

		case "3\n":
			return "LIST"

		case "4\n", ":q\n":
			return "STOP"

		default:
			fmt.Println("That option doesn't exist")
			continue
		}
	}
}
