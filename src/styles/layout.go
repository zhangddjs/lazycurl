package styles

import "github.com/charmbracelet/lipgloss"

// size
var (
	// filemanager
	FileManagerW = 20
	FileManagerH = 17

	// filemanager
	BufferManagerW = 20
	BufferManagerH = 34

	// logo
	LogoW = 20
	LogoH = 1

	// Method
	MethodW = 6
	MethosH = 1

	//Url
	UrlW = 44
	UrlH = 1

	//Req Body
	ReqBodyW = 52
	ReqBodyH = 16

	//Req Body
	RespBodyW = 52
	RespBodyH = 16
)

// color
var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

var (
	TextAreaStyle = lipgloss.NewStyle().
			Width(ReqBodyW).
			Height(ReqBodyH).
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder())

	FocusedTextAreaStyle = lipgloss.NewStyle().
				Width(ReqBodyW).
				Height(ReqBodyH).
				Align(lipgloss.Left, lipgloss.Top).
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

	BufferManagerStyle = lipgloss.NewStyle().
				Width(BufferManagerW).
				Height(BufferManagerH).
				Align(lipgloss.Left, lipgloss.Top).
				BorderStyle(lipgloss.NormalBorder())

	FocusedBufferManagerStyle = lipgloss.NewStyle().
					Width(BufferManagerW).
					Height(BufferManagerH).
					Align(lipgloss.Left, lipgloss.Top).
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(highlight)

	RespBodyStyle = lipgloss.NewStyle().
			Width(RespBodyW).
			Height(RespBodyH).
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder())

	FocusedRespBodyStyle = lipgloss.NewStyle().
				Width(RespBodyW).
				Height(RespBodyH).
				Align(lipgloss.Left, lipgloss.Top).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(highlight)

	ReqStyle = lipgloss.NewStyle().
			Width(ReqBodyW).
			Height(ReqBodyH).
			Align(lipgloss.Left, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder())

	FocusedReqStyle = lipgloss.NewStyle().
			Width(ReqBodyW).
			Height(ReqBodyH).
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
