package tui

import (
	"fmt"
	"sgbd-1/src/storage"

	tea "github.com/charmbracelet/bubbletea"
)

type option int

const (
	none option = iota
	scan
	seek
	insert
	del
)

func (opt option) ToString() string {
	switch opt {
	case scan:
		return "Scanear"
	case seek:
		return "Buscar por conteúdo"
	case insert:
		return "Inserir novo documento"
	case del:
		return "Remover documento"
	default:
		return ""
	}
}

func OptionFrom(v int) option {
	switch v {
	case 1:
		return scan
	case 2:
		return seek
	case 3:
		return insert
	case 4:
		return del
	}

	return none
}

type model struct {
	options     []option
	currentSate option
	cursorPos   int
	storage     storage.Storage
}

func InitialModel() model {
	return model{
		options:     []option{scan, seek, insert, del},
		currentSate: scan,
		storage:     storage.GetStorage(),
	}
}
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "w":
			if m.cursorPos > 0 {
				m.cursorPos--
			}

		case "down", "s":
			if m.cursorPos < len(m.options)-1 {
				m.cursorPos++
			}

		case "enter", " ":
			m.currentSate = OptionFrom(m.cursorPos + 1)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "O que deseja fazer?\n"

	for i, choice := range m.options {

		cursor := " "
		if m.cursorPos == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.ToString())
	}

	s += "\nPressione q para sair.\n"

	return s
}
