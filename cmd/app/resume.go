package resume

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// Style definitions.
var (
	headerStyle = lipgloss.NewStyle().
		//			BorderStyle(lipgloss.NormalBorder()).
		//BorderForeground(lipgloss.Color("#874BFD")).
		Padding(0, 1)

	activeTabStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#874BFD")).
			Padding(0, 1)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Padding(0, 1)

	splashStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
		//			Background(lipgloss.Color("#874BFD")).
		// BorderStyle(lipgloss.DoubleBorder()).
		// BorderForeground(lipgloss.Color("#FAFAFA")).
		// Padding(1, 0).
		Bold(true)

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()

	// Status Bar.

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	// encodingStyle = statusNugget.
	// 		Background(lipgloss.Color("#A550DF")).
	// 		Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Background(lipgloss.Color("#6124DF"))

	//spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type model struct {
	showingSplash bool
	currentView   int
	viewports     []viewport.Model
	viewport      viewport.Model
	width         int
	height        int
	Ticks         int
	ready         bool
}

func InitialModel() model {
	return model{
		showingSplash: true,
		currentView:   0,
		viewports:     make([]viewport.Model, 3),
	}
}

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(2*time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() tea.Cmd {

	return tick()
}

func (m model) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var homeContent = homePage()
	// var sobreContent = sobrePage()
	// var contatoContent = contatoPage()
	// Navigation.
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tickMsg:
		if m.Ticks == 0 {
			m.showingSplash = false
		}
		m.Ticks--
		return m, tick()
	case tea.KeyMsg:
		if m.showingSplash {
			return m, nil
		}
		switch strings.ToLower(msg.String()) {
		case "ctrl+c", "q":

			return m, tea.Quit
		case "p":
			m.currentView = 0
		case "s":
			m.currentView = 1
		case "c":
			m.currentView = 2
		case "right", "l":
			m.currentView = (m.currentView + 1) % len(m.viewports)
		case "left", "h":
			m.currentView = (m.currentView - 1 + len(m.viewports)) % len(m.viewports)
		case "up", "down", "k", "j":
			m.viewports[m.currentView], cmd = m.viewports[m.currentView].Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		//headerHeight := 3
		//footerHeight := 1
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())

		viewportHeight := m.height - headerHeight - footerHeight
		verticalMarginHeight := m.height + headerHeight + footerHeight

		//if m.viewports[0].Height == 0 {
		if !m.ready {

			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			//m.viewport.HighPerformanceRendering = useHighPerformanceRenderer

			for i, content := range []string{homeContent, sobreContent, contatoContent} {
				m.viewports[i] = viewport.New(m.width, viewportHeight)
				m.viewports[i].Style = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("#874BFD"))

				renderedContent, _ := glamour.Render(content, "dark")
				m.viewports[i].SetContent(renderedContent)
			}
			m.ready = true
		} else {
			//for i := range m.viewports {
			//m.viewports[i].Width = m.width
			//m.viewports[i].Height = viewportHeight
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
			//}
		}

	}

	return m, nil
}

func (m model) View() string {
	splashContent := "CRISTIANO LEMES"

	if m.showingSplash {

		splashText := fmt.Sprintf("%s", splashContent)
		style := splashStyle
		return style.Render(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, splashText))
	}

	tabs := []string{"p Principal", "s Sobre", "c Contato"}
	renderedTabs := make([]string, len(tabs))

	for i, tab := range tabs {
		if i == m.currentView {
			renderedTabs[i] = activeTabStyle.Render(tab)
		} else {
			renderedTabs[i] = inactiveTabStyle.Render(tab)
		}
	}

	header := headerStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, renderedTabs...))

	// Status bar

	w := lipgloss.Width

	statusKey := statusStyle.Render("RUNNING")
	//encoding := encodingStyle.Render("UTF-8")
	fishCake := fishCakeStyle.Render("⚡ Cris.Run")
	statusVal := statusText.
		//Width(m.width - w(statusKey) - w(encoding) - w(fishCake)).
		Width(m.width - w(statusKey) - w(fishCake)).
		Render("Pressione 'q' para sair.")

	bar := statusBarStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		//encoding,
		fishCake,
	))

	//doc.WriteString(statusBarStyle.Width(width).Render(bar))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.viewports[m.currentView].View(),
		bar,
	)
}

// func main() {
// 	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Erro ao executar programa: %v", err)
// 		os.Exit(1)
// 	}
// }
