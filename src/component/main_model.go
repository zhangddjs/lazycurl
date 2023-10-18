package component

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	fm "github.com/zhangddjs/lazycurl/component/filemanager"
	hm "github.com/zhangddjs/lazycurl/component/httpmethod"
	rq "github.com/zhangddjs/lazycurl/component/request"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/model"
	"github.com/zhangddjs/lazycurl/styles"
)

type sessionState uint

var (
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

const (
	fmView sessionState = iota
	bmView
	methodView
	urlView
	reqView
	respBodyView
)

const (
	MODEL_CNT   = 6
	defaultTime = time.Minute
)

type mainModel struct {
	state          sessionState
	method         hm.Model
	url            vp.Model
	req            rq.Model
	reqBody        rq.BodyModel
	respBody       vp.Model
	filemanager    fm.Model
	bufmanager     fm.BufModel
	analyzer       fm.AnalyzerModel
	activeFileNode *model.FileNode
	activeCurl     *model.Curl
	index          int
}

func NewModel(timeout time.Duration) mainModel {
	m := mainModel{state: fmView}
	m.filemanager = fm.New()
	m.bufmanager = fm.NewBM()
	m.analyzer = fm.NewAnalyzer()
	m.url = vp.New(styles.UrlW, styles.UrlH, "https://www.youtube.com/watch?v=\n_F0-q1jeReY&list=PL-3c1Yp7oGX8MLyYp1-uFq8RMGRQ00whV&index=122&ab_channel=supershigi")
	m.req = rq.New(styles.ReqBodyW, styles.ReqBodyH)
	m.reqBody = rq.NewBodyModel(styles.ReqBodyW, styles.ReqBodyH)
	m.respBody = vp.New(styles.RespBodyW, styles.RespBodyH, "{\n\taaa:bbb\n}")
	return m
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = m.handleKey(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		cmd = m.handleWindowSize(msg)
		return m, cmd
	case model.SuccessMsg:
		cmd = m.handleFmSuccess(msg)
		return m, cmd
	case model.AnalyzeMsg:
		cmd = m.handleAnalyze(msg)
		return m, cmd

	case hm.SwitchMethodMsg:
		m.activeCurl.SetMethod(m.method.GetMethod())
		m.activeFileNode.SetBuffer(m.activeCurl.BuildCurlCmd())
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s = strings.Builder{}
	model := m.currentFocusedModel()

	fm := m.filemanager.Render(m.isActive(fmView))
	bm := m.bufmanager.Render(m.isActive(bmView))
	logo := styles.LogoStyle.Render()
	method := m.method.Render(m.isActive(methodView))
	url := m.url.RenderUrl(m.isActive(urlView))
	req := m.req.Render(m.isActive(reqView))
	reqBody := m.reqBody.Render(m.isActive(respBodyView))
	respBody := m.respBody.RenderRespBody(m.isActive(respBodyView))
	_ = respBody

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, logo, method, url))
	s.WriteString("\n")
	fmField := lipgloss.JoinVertical(lipgloss.Left, fm, bm)
	txtArea := lipgloss.JoinVertical(lipgloss.Left, req, reqBody)
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, fmField, txtArea))
	s.WriteString(helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model)))
	return s.String()
}

func (m *mainModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.String() {
	case "ctrl+c":
		return tea.Quit
	case "tab":
		m.SwitchToNextModel()
		// TODO: ctrl+s save to file
	}
	switch m.state {
	// update whichever model is focused
	case fmView:
		m.filemanager, cmd = m.filemanager.Update(msg)
		cmds = append(cmds, cmd)
	case bmView:
		m.bufmanager, cmd = m.bufmanager.Update(msg)
		cmds = append(cmds, cmd)
	case urlView:
		m.url, cmd = m.url.Update(msg)
		cmds = append(cmds, cmd)
	case methodView:
		m.method, cmd = m.method.Update(msg)
		cmds = append(cmds, cmd)
	case reqView:
		m.req, cmd = m.req.Update(msg)
		cmds = append(cmds, cmd)
	case respBodyView:
		m.respBody, cmd = m.respBody.Update(msg)
		cmds = append(cmds, cmd)
		m.reqBody, cmd = m.reqBody.Update(msg) // TODO: need remove this
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (m *mainModel) handleWindowSize(msg tea.WindowSizeMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.url, cmd = m.url.Update(msg)
	cmds = append(cmds, cmd)
	m.req, cmd = m.req.Update(msg)
	cmds = append(cmds, cmd)
	m.respBody, cmd = m.respBody.Update(msg)
	cmds = append(cmds, cmd)
	m.reqBody, cmd = m.reqBody.Update(msg) //TODO: need remove this
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

// handleFmSuccess
// 1.ReadFileSuccess -- buff, analyzer
func (m *mainModel) handleFmSuccess(msg model.SuccessMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.Type {
	case model.ReadFileSuccess:
		data := msg.Data.(model.ReadFileSuccessData)
		m.bufmanager, cmd = m.bufmanager.Update(msg)
		m.activeFileNode = data.Item
		cmds = append(cmds, cmd, model.Analyze(data.Item))
	case model.OpenBufferSuccess:
		data := msg.Data.(model.OpenBufferSuccessData)
		m.activeFileNode = data.Item
		// TODO: need file manager know about this and expand dirs, update his cursor
		cmds = append(cmds, cmd, model.Analyze(data.Item))
	case model.AnalyzeSuccess:
		data := msg.Data.(model.AnalyzeSuccessData)
		m.activeCurl = data.Curl
		m.method.SetMethod(strings.ToUpper(data.Curl.GetMethod()))
		m.url.SetContent(data.Curl.GetUrl())
		m.req, cmd = m.req.Update(msg)
		m.reqBody.SetBody(data.Curl.GetBody())
		cmds = append(cmds, cmd)
		// TODO: render
	}
	return tea.Batch(cmds...)
}

func (m *mainModel) handleAnalyze(msg model.AnalyzeMsg) tea.Cmd {
	var cmd tea.Cmd
	m.analyzer, cmd = m.analyzer.Update(msg)
	return cmd
}

func (m mainModel) currentFocusedModel() string {
	return "spinner"
}

func (m mainModel) isActive(state sessionState) bool {
	return m.state == state
}

func (m *mainModel) SwitchToNextModel() {
	m.state = (m.state + 1) % MODEL_CNT
}
