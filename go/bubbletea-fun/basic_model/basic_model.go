package basic_model

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var greetings = []string{
	"hello",
	"hola",
	"bonjour",
}

type model struct {
	greetings []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	nItems := int64(len(greetings))
	randIndex := time.Now().Unix() % nItems
	return m.greetings[randIndex]
}

func Run() {
	p := tea.NewProgram(model{greetings: greetings})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
