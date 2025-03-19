package basic_keymap

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// for each number 1..5, show a new styling

type model struct {
	currentStyle lipgloss.Style
	keymap       KeyMap
	styles       Styles
}

type KeyMap struct {
	One   key.Binding
	Two   key.Binding
	Three key.Binding
	Four  key.Binding
	Five  key.Binding
	Quit  key.Binding
}

var DefaultKeyMap = KeyMap{
	One:   key.NewBinding(key.WithKeys("1")),
	Two:   key.NewBinding(key.WithKeys("2")),
	Three: key.NewBinding(key.WithKeys("3")),
	Four:  key.NewBinding(key.WithKeys("4")),
	Five:  key.NewBinding(key.WithKeys("5")),
	Quit:  key.NewBinding(key.WithKeys("q", "exc", "ctrl+c", "ctrl+d")),
}

type Styles struct {
	One   lipgloss.Style // placeholder for lipgloss function
	Two   lipgloss.Style
	Three lipgloss.Style
	Four  lipgloss.Style
	Five  lipgloss.Style
}

var DefaultStyles = Styles{
	One: lipgloss.NewStyle().
		Blink(true).
		Foreground(lipgloss.Color("31")),
	Two: lipgloss.NewStyle().
		Align(lipgloss.Center).
		Bold(true),
	Three: lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("10")).
		BorderForeground(lipgloss.Color("89")),
	Four: lipgloss.NewStyle().
		Reverse(true),
	Five: lipgloss.NewStyle().
		Underline(true).
		Padding(1).
		Margin(4).
		MarginBackground(lipgloss.Color("3")),
}

func (m model) Init() tea.Cmd { return nil }

func (m model) View() string {
	return m.currentStyle.Render("what is happening!")
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.One):
			m.currentStyle = m.styles.One
		case key.Matches(msg, m.keymap.Two):
			m.currentStyle = m.styles.Two
		case key.Matches(msg, m.keymap.Three):
			m.currentStyle = m.styles.Three
		case key.Matches(msg, m.keymap.Four):
			m.currentStyle = m.styles.Four
		case key.Matches(msg, m.keymap.Five):
			m.currentStyle = m.styles.Five
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func Run() {
	p := tea.NewProgram(model{
		keymap: DefaultKeyMap,
		styles: DefaultStyles,
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
