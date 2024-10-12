package pages

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/cslemes/bbbb/cmd/config"
	"github.com/cslemes/bbbb/cmd/utils"
)

type model struct {
	sessInfo             ssh.Session
	showingSplash        bool
	currentView          int
	viewports            []viewport.Model
	viewport             viewport.Model
	width                int
	height               int
	Ticks                int
	ready                bool
	selectedItem         int
	listItems            []string
	selectedContent      viewport.Model
	focusedOnList        bool
	viewportHeight       int
	verticalMarginHeight int
	viewportWidth        int
	blogPageOpen         bool
	//
	color               lipgloss.AdaptiveColor
	activeSelectStyle   lipgloss.Style
	inactiveSelectStyle lipgloss.Style
	activeTabBorder     lipgloss.Border
	tabBorder           lipgloss.Border
	tab                 lipgloss.Style
	tabGap              lipgloss.Style
	activeTab           lipgloss.Style
	splashStyle         lipgloss.Style
	titleStyle          lipgloss.Style
	infoStyle           lipgloss.Style
	defaultStyle        lipgloss.Style
	statusNugget        lipgloss.Style
	statusStyle         lipgloss.Style
	fishCakeStyle       lipgloss.Style
	statusBarStyle      lipgloss.Style
	statusText          lipgloss.Style
	blends              []color.Color
}

func InitialModel(sess ssh.Session) model {

	listItems, err := utils.LoadFilesFromDir("posts")
	if err != nil {
		fmt.Println("Error loading files from posts:", err)
		listItems = []string{"Error loading files"}
	}

	return model{
		sessInfo:        sess,
		showingSplash:   config.AppConfig().Navigation.ShowingSplash,
		currentView:     0,
		viewports:       make([]viewport.Model, 5),
		listItems:       listItems,
		selectedContent: viewport.New(0, 0),
		focusedOnList:   true,
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

	tabs := []string{"p Principal", "a About", "c Contact", "b Blog"}
	renderedTabs := make([]string, len(tabs))

	for i, menu := range tabs {
		if i == m.currentView {
			renderedTabs[i] = m.configTheme().activeTab.Render(menu)
		} else {
			renderedTabs[i] = m.configTheme().tab.Render(menu)
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	gap := m.configTheme().tabGap.Render(strings.Repeat(" ", max(0, lipgloss.Width(row)-2)))

	return lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

}

func (m model) footerView() string {
	w := lipgloss.Width

	statusKey := m.configTheme().statusStyle.Render("RUNNING")
	fishCake := m.configTheme().fishCakeStyle.Render("⚡ Cris.Run")
	statusVal := m.configTheme().statusText.
		Width(m.width - w(statusKey) - w(fishCake)).
		Render("Pressione 'q' para sair.")

	if m.currentView == 4 {
		statusVal = m.configTheme().statusText.
			Width(m.width - w(statusKey) - w(fishCake)).
			Render("Pressione 'q' para sair. Pressione 'Esc' para sair do modo letura.")

	}

	return m.statusBarStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		fishCake,
	))

}

func (m model) listView() string {

	var listBuilder strings.Builder
	for i, item := range m.listItems {
		if i == m.selectedItem {
			listBuilder.WriteString(m.configTheme().activeSelectStyle.Render("> " + item))
		} else {
			listBuilder.WriteString(m.configTheme().inactiveSelectStyle.Render("  " + item))
		}
		listBuilder.WriteString("\n")
	}

	//return listBuilder.String()
	return lipgloss.NewStyle().
		Render(listBuilder.String())

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "p":
			m.currentView = 0
		case "a":
			m.currentView = 1
		case "c":
			m.currentView = 2
		case "b":
			if m.blogPageOpen || !m.focusedOnList {
				m.currentView = 4
			} else {
				m.currentView = 3
			}
		case "right", "l":
			m.currentView = (m.currentView + 1) % (len(m.viewports) - 1)

		case "left", "h":
			m.currentView = (m.currentView - 1 + len(m.viewports)) % len(m.viewports)

		case "up", "k":
			if m.currentView == 3 {
				if m.focusedOnList {
					if m.selectedItem > 0 {
						m.selectedItem--
					}
				}
			} else {
				m.viewports[m.currentView], cmd = m.viewports[m.currentView].Update(msg)
			}

		case "down", "j":
			if m.currentView == 3 {
				if m.focusedOnList {
					if m.selectedItem < len(m.listItems)-1 {
						m.selectedItem++
					}
				}
			} else {
				m.viewports[m.currentView], cmd = m.viewports[m.currentView].Update(msg)
			}

		case "enter":
			if m.currentView == 3 {
				if m.focusedOnList {
					filename := m.listItems[m.selectedItem]
					content, err := utils.ReadFileContent("posts/" + filename + ".md")
					if err != nil {
						m.selectedContent.SetContent("Error reading file: " + err.Error())
					} else {
						renderedContent, _ := glamour.Render(content, "dark")
						m.viewports[4] = viewport.New(m.width, m.viewportHeight)
						m.viewports[4].Style = m.configTheme().defaultStyle
						m.viewports[4].SetContent(renderedContent)
						m.currentView = 4
						m.focusedOnList = false
						m.blogPageOpen = true
					}
				}
			}
		case "esc":
			m.currentView = 3
			m.focusedOnList = true

			// case "up", "down", "k", "j":
			// 	m.viewports[m.currentView], cmd = m.viewports[m.currentView].Update(msg)

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())

		m.viewportHeight = m.height - headerHeight - footerHeight
		m.verticalMarginHeight = m.height + headerHeight + footerHeight
		m.viewportWidth = m.width

		if !m.ready {

			m.viewport = viewport.New(msg.Width, msg.Height-m.verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.navigation().View())

			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - m.verticalMarginHeight
		}

	}
	//m.viewports[m.currentView], cmd = m.viewports[m.currentView].Update(msg)
	return m, cmd
}

func (m model) View() string {

	setSplash := utils.GetTerminalColorSupport(m.sessInfo)

	splashContent := splashContent(setSplash)

	if m.showingSplash {
		splashText := splashContent
		style := m.splashStyle
		return style.Render(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, splashText))
	}

	m.blogPage()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerView(),
		m.viewports[m.currentView].View(),

		m.footerView(),
	)

}

func (m model) navigation() tea.Model {

	var homeContent = homePage()
	var sobreContent = sobrePage()
	var contatoContent = contatoPage()

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.width),
	)

	for i, content := range []string{homeContent, sobreContent, contatoContent} {

		m.viewports[i] = viewport.New(m.width, m.viewportHeight)
		m.viewports[i].Style = m.configTheme().defaultStyle

		renderedContent, _ := renderer.Render(content)
		m.viewports[i].SetContent(renderedContent)

	}
	return m
}

func (m model) blogPage() tea.Model {

	m.viewports[3] = viewport.New(m.width, m.viewportHeight)
	m.viewports[3].Style = m.configTheme().defaultStyle
	m.viewports[3].SetContent(m.listView())

	return m

}
