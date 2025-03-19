package accordian

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Setup */

type model struct {
	keymap         keymap
	accordianItems accordianItems
}

type keymap struct {
	quit key.Binding
	up   key.Binding
	down key.Binding
	sel  key.Binding
}

type accordianItems struct {
	items map[string][]string
	focus string
}

func (a accordianItems) View() string {
	var results string
	// Return all the keys if there is no focus
	if a.focus == "" {
		for key := range a.items {
			results += "\n" + key
		}
		return results
	}
	// Return the keys and values for the focued item
	for key, values := range a.items {
		results += "\n" + key
		if a.focus == key {
			for _, v := range values {
				results += "\n  " + v
			}
		}
	}

	return results
}

var selectedStyle = lipgloss.NewStyle().Bold(true)

/* Init Functions */

func newModel() model {
	return model{
		keymap:         newKeymap(),
		accordianItems: newAccordianItems(),
	}
}

func newKeymap() keymap {
	return keymap{
		quit: key.NewBinding(key.WithKeys("q")),
		up:   key.NewBinding(key.WithKeys("k")),
		down: key.NewBinding(key.WithKeys("j")),
		sel:  key.NewBinding(key.WithKeys(" ")),
	}
}

func newAccordianItems() accordianItems {
	focus := ""
	results := make(map[string][]string)

	results["a"] = []string{"1", "2"}
	results["b"] = []string{"3", "4"}
	results["c"] = []string{"5", "6"}

	return accordianItems{
		items: results,
		focus: focus,
	}
}

/* Body Functions */

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.up):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.down):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.sel):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return m.accordianItems.View()
}

func Run() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
