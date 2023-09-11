package styles

import "github.com/charmbracelet/lipgloss"

// size
var (
	// filemanager
	FileManagerW = 20
	FileManagerH = 34

	// logo
	LogoW = 20
	LogoH = 1

	// Method
	MethodW = 6
	MethosH = 1

	//Url
	UrlW = 44
	UrlH = 1
)

// color
var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

var (
	ModelStyle = lipgloss.NewStyle().
			Width(52).
			Height(16).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder())

	FocusedModelStyle = lipgloss.NewStyle().
				Width(52).
				Height(16).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(highlight)

	LogoStyle = lipgloss.NewStyle().
			Width(LogoW).
			Height(LogoH).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			SetString("LazyCurl v0.0.1")

	MethodStyle = lipgloss.NewStyle().
			Width(MethodW).
			Height(MethosH).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder())

	FocusedMethodStyle = lipgloss.NewStyle().
				Width(MethodW).
				Height(MethosH).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(highlight)

	UrlStyle = lipgloss.NewStyle().
			Width(UrlW).
			Height(UrlH).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder())

	FocusedUrlStyle = lipgloss.NewStyle().
			Width(UrlW).
			Height(UrlH).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(highlight)

	FileManagerStyle = lipgloss.NewStyle().
				Width(FileManagerW).
				Height(FileManagerH).
				Align(lipgloss.Left, lipgloss.Top).
				BorderStyle(lipgloss.NormalBorder())

	FocusedFileManagerStyle = lipgloss.NewStyle().
				Width(FileManagerW).
				Height(FileManagerH).
				Align(lipgloss.Left, lipgloss.Top).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(highlight)
)

var (
	TitleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)
