package paragraphs

/*
Interesting things:
- Use m.ready to setup the hieght / width of the viewport
- Use a bubble
  - Add the bubble.Model to our model
  - Initialize it with bubble.New()
  - Pass update commands to it, and keep track of cmds it gives out
  - Batch response commands in updates returen
  - Render the bubble in our return
  - Once the bubble is set, you can access its properties
  - State! okay so the basic idea is to create a model state
    then depending on the state, call the submodel's update
    and batch those to the cmd
*/

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Idea: Display paragraphs of text next to eachother so it looks cool

/* Setup */
type sessionState uint

const (
	topMargin                    = 1
	viewportDivisor              = 2
	bluebirdPath                 = "./paragraphs/bluebird.txt"
	geniusPath                   = "./paragraphs/genius.txt"
	geniusState     sessionState = iota
	bluebirdState
)

type poem struct {
	text     string
	title    string
	viewport viewport.Model
}

func (p *poem) Ready(m model, msg tea.WindowSizeMsg) {
	p.viewport = viewport.New(msg.Width/viewportDivisor, msg.Height-topMargin)
	p.viewport.SetContent(p.text)
	p.viewport.YPosition = topMargin
	p.viewport.Style = m.styles.paragraph
}

func (p *poem) Resize(msg tea.WindowSizeMsg) {
	p.viewport.Height = msg.Height - topMargin
	p.viewport.Width = msg.Width / viewportDivisor
}

func (p *poem) View(m model) string {
	return fmt.Sprintf("\n\n%s\n%s", m.styles.title.Render(p.title), p.viewport.View())
}

type model struct {
	keymap   keymap
	styles   styles
	genius   poem
	bluebird poem
	ready    bool
	state    sessionState
}

type keymap struct {
	quit       key.Binding
	changeView key.Binding
}

type styles struct {
	paragraph lipgloss.Style
	title     lipgloss.Style
}

var geniusPoem = poem{
	title: "genius of the crowd",
	text:  loadText(geniusPath),
}

var bluebirdPoem = poem{
	title: "bluebird",
	text:  loadText(bluebirdPath),
}

var defaultModel = model{
	keymap:   defaultKeymap,
	styles:   defaultStyles,
	genius:   geniusPoem,
	bluebird: bluebirdPoem,
	state:    geniusState,
}

var defaultKeymap = keymap{
	quit:       key.NewBinding(key.WithKeys("q", "esc", "ctrl+c", "ctrl+d")),
	changeView: key.NewBinding(key.WithKeys("tab")),
}

var defaultStyles = styles{
	paragraph: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Margin(1).Padding(1, 3),
	title:     lipgloss.NewStyle().Bold(true).Margin(0, 3),
}

/* Run */
func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.changeView):
			if m.state == geniusState {
				m.state = bluebirdState
			} else {
				m.state = geniusState
			}
		}
		switch m.state {
		case geniusState:
			m.genius.viewport, cmd = m.genius.viewport.Update(msg)
		case bluebirdState:
			m.bluebird.viewport, cmd = m.bluebird.viewport.Update(msg)
		}
	case tea.WindowSizeMsg:
		if !m.ready {
			// Ready
			m.genius.Ready(m, msg)
			m.bluebird.Ready(m, msg)
			m.ready = true
		} else {
			// Resize
			m.genius.Resize(msg)
			m.bluebird.Resize(msg)
		}
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, m.genius.View(m), m.bluebird.View(m))
}

func Run() {
	p := tea.NewProgram(
		defaultModel,
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

/* Helpers */
func loadText(path string) string {
	fmt.Println(os.Getwd())
	file, err := os.ReadFile(path)
	if err != nil {
		panic("Unable to load file")
	}

	return string(file)
}
