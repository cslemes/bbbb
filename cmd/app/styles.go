package pages

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/cslemes/bbbb/cmd/config"
	"github.com/muesli/gamut"
)

func (m model) configTheme() model {

	Config := config.AppConfig()

	selectedColor := Config.Theme.Color
	switch selectedColor {
	case "subtle":
		m.color = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	case "highlight":
		m.color = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	case "special":
		m.color = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	case "bubble":
		m.color = lipgloss.AdaptiveColor{Light: "#AD2174", Dark: "#FF7692"}
	}
	m.blends = gamut.Blends(lipgloss.Color("#F25D94"), lipgloss.Color("#EDFF82"), 50)

	m.activeSelectStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#874BFD")).
		Padding(0, 1)

	m.inactiveSelectStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Padding(0, 1)

	m.activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "║",
		TopLeft:     "╭",
		TopRight:    "╖",
		BottomLeft:  "┘",
		BottomRight: "╙",
	}

	m.tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	m.tab = lipgloss.NewStyle().
		Border(m.tabBorder, true).
		BorderForeground(m.color).
		Padding(0, 1)

	m.activeTab = m.tab.Border(m.activeTabBorder, true)

	m.tabGap = m.tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	m.splashStyle = lipgloss.NewStyle().
		Foreground(m.color).
		//			Background(lipgloss.Color("#874BFD")).
		// BorderStyle(lipgloss.DoubleBorder()).
		// BorderForeground(lipgloss.Color("#FAFAFA")).
		// Padding(1, 0).
		Bold(true)

	m.titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	m.infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return m.titleStyle.BorderStyle(b)
	}()

	m.defaultStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	// Status Bar.

	m.statusNugget = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	m.statusBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	m.statusStyle = lipgloss.NewStyle().
		Inherit(m.statusBarStyle).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)

	m.statusText = lipgloss.NewStyle().Inherit(m.statusBarStyle)

	m.fishCakeStyle = m.statusNugget.Background(lipgloss.Color("#6124DF"))

	//spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return m
}
