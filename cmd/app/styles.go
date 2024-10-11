package pages

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/cslemes/bbbb/cmd/config"
)

// Style definitions.
// type Theme struct {
// 	color               lipgloss.AdaptiveColor
// 	activeSelectStyle   lipgloss.Style
// 	inactiveSelectStyle lipgloss.Style
// 	activeTabBorder     lipgloss.Border
// 	tabBorder           lipgloss.Border
// 	tab                 lipgloss.Style
// 	tabGap              lipgloss.Style
// 	activeTab           lipgloss.Style
// 	splashStyle         lipgloss.Style
// 	titleStyle          lipgloss.Style
// 	infoStyle           lipgloss.Style
// 	defaultStyle        lipgloss.Style
// 	statusNugget        lipgloss.Style
// 	statusStyle         lipgloss.Style
// 	fishCakeStyle       lipgloss.Style
// 	statusBarStyle      lipgloss.Style
// 	statusText          lipgloss.Style
// }

func (t model) configTheme() model {

	Config := config.AppConfig()

	selectedColor := Config.Theme.Color
	switch selectedColor {
	case "subtle":
		t.color = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	case "highlight":
		t.color = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	case "special":
		t.color = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	}
	//blends := gamut.Blends(lipgloss.Color("#F25D94"), lipgloss.Color("#EDFF82"), 50)

	t.activeSelectStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#874BFD")).
		Padding(0, 1)

	t.inactiveSelectStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Padding(0, 1)

	t.activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "║",
		TopLeft:     "╭",
		TopRight:    "╖",
		BottomLeft:  "┘",
		BottomRight: "╙",
	}

	t.tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	t.tab = lipgloss.NewStyle().
		Border(t.tabBorder, true).
		BorderForeground(t.color).
		Padding(0, 1)

	t.activeTab = t.tab.Border(t.activeTabBorder, true)

	t.tabGap = t.tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	t.splashStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		//			Background(lipgloss.Color("#874BFD")).
		// BorderStyle(lipgloss.DoubleBorder()).
		// BorderForeground(lipgloss.Color("#FAFAFA")).
		// Padding(1, 0).
		Bold(true)

	t.titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	t.infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return t.titleStyle.BorderStyle(b)
	}()

	t.defaultStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	// Status Bar.

	t.statusNugget = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	t.statusBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	t.statusStyle = lipgloss.NewStyle().
		Inherit(t.statusBarStyle).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)

	t.statusText = lipgloss.NewStyle().Inherit(t.statusBarStyle)

	t.fishCakeStyle = t.statusNugget.Background(lipgloss.Color("#6124DF"))

	//spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return t
}
