package tui

import (
	"bufio"
	"fmt"
	"os"
	"sgbd-1/src/doc"
	"sgbd-1/src/page"
	"sgbd-1/src/storage"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	itemStyle      = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false).Width(60)
	containerStyle = lipgloss.NewStyle().Padding(0, 1).AlignHorizontal(lipgloss.Left).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#552b8f"))

	pageHeaderInfoStyle   = lipgloss.NewStyle().Border(lipgloss.ThickBorder(), false, false, true, false).Width(52).PaddingLeft(2)
	docsInfoListContainer = lipgloss.NewStyle().Inherit(containerStyle).BorderForeground(lipgloss.Color("#ffffff")).Width(50).Padding(0)

	errorStyle = lipgloss.NewStyle().Padding(0, 1).AlignHorizontal(lipgloss.Left).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#ff0000"))
)

type TUIDocs struct {
	currentOption state
	storage       storage.Storage
}

func NewTUI() *TUIDocs {
	return &TUIDocs{quit, storage.GetStorage()}
}

func (tui *TUIDocs) RenderPages() {
	var free []string
	var used []string

	space := strings.Repeat(" ", 20)
	free = append(free, headerStyle.Render(fmt.Sprintf("%sPáginas livres%s\n", space, space)))

	curr := tui.storage.FreePages().Head()

	for curr != nil {
		free = append(free, tui.PageTile(curr.Page()))
		curr = curr.Next()
	}

	freePages := lipgloss.JoinVertical(lipgloss.Left, free...)

	used = append(used, headerStyle.Render(fmt.Sprintf("%sPáginas usadas%s\n", space, space)))

	curr = tui.storage.UsedPages().Head()

	for curr != nil {
		used = append(used, tui.PageTile(curr.Page()))
		curr = curr.Next()
	}

	usedPages := lipgloss.JoinVertical(lipgloss.Left, used...)

	usedPages = containerStyle.Render(usedPages)
	freePages = containerStyle.Render(freePages)

	pages := lipgloss.JoinHorizontal(lipgloss.Left, usedPages, freePages)

	fmt.Println(containerStyle.Render(pages))

}

func (tui *TUIDocs) PageInfoHeader(page page.Page) string {
	return pageHeaderInfoStyle.Render(fmt.Sprintf("ID -> %*s | Tamanho %d", 2, strconv.Itoa(page.ID()), page.Size()))
}

func (tui *TUIDocs) PageTile(page page.Page) string {
	pageInfo := tui.PageInfoHeader(page)
	docsList := tui.DocsContainer(page.GetDocuments())

	return lipgloss.JoinVertical(lipgloss.Left, pageInfo, docsList)
}

func (tui *TUIDocs) DocInfo(doc *doc.Document) string {
	return fmt.Sprintf("PID %*s | Pos %d | Tamanho %d | Conteúdo %s", 2, strconv.Itoa(doc.PageID), doc.Position, doc.Length, doc.Content)
}

func (tui *TUIDocs) DocsContainer(docs []*doc.Document) string {
	var lines []string

	if len(docs) == 0 {
		return ""
	}

	for _, doc := range docs {
		lines = append(lines, tui.DocInfo(doc))
	}

	docsInfoList := lipgloss.JoinVertical(lipgloss.Left, lines...)

	return docsInfoListContainer.Render(docsInfoList)
}

func (tui *TUIDocs) RenderMenu() {
	var lines []string
	states := []state{scan, seek, add, del, quit}
	command := []string{"scan", "seek (param)", "add (param)", "del (param)", "quit"}

	space := strings.Repeat(" ", 20)
	lines = append(lines, headerStyle.Render(fmt.Sprintf("%sMenu%s\n", space, space)))
	for i := range states {
		lines = append(lines, itemStyle.Render(fmt.Sprintf("[%-*s] %v", 15, command[i], states[i])))
	}

	menu := lipgloss.JoinVertical(lipgloss.Left, lines...)
	fmt.Println(containerStyle.Render(menu))
}

func (tui *TUIDocs) parseInput(cmd string) (state, []byte, error) {
	if strings.Contains(cmd, "scan") {
		return scan, nil, nil
	}
	if strings.Contains(cmd, "seek") {
		query := strings.Split(strings.Trim(cmd, " "), " ")
		if len(query) < 2 {
			return fail, nil, fmt.Errorf("seek espera um argumento do conteúdo do documento")
		}
		param := query[1]
		return seek, []byte(param), nil
	}
	if strings.Contains(cmd, "del") {
		query := strings.Split(strings.Trim(cmd, " "), " ")
		if len(query) < 2 {
			return fail, nil, fmt.Errorf("del espera um argumento do conteúdo do documento")
		}
		param := query[1][:len(query)]
		return del, []byte(param), nil
	}

	if strings.Contains(cmd, "add") {
		query := strings.Split(strings.Trim(cmd, " "), " ")
		if len(query) < 2 {
			return fail, nil, fmt.Errorf("add espera um argumento do conteúdo do documento")
		}
		param := query[1][:len(query)]
		return add, []byte(param), nil
	}

	if strings.Contains(cmd, "quit") {
		return quit, nil, nil
	}

	return fail, nil, fmt.Errorf("entrada não reconhecida")

}

func (tui *TUIDocs) Run() {
	scanner := bufio.NewReader(os.Stdin)

	tui.RenderMenu()

	for {

		cmd, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler stdin.", err.Error())
		}

		state, param, e := tui.parseInput(cmd)
		if err != nil {
			fmt.Println(errorStyle.Render(e.Error()))
		}

		switch state {
		case scan:
			fmt.Println(state)

		case seek:
			fmt.Println(state, param)
		case add:
			fmt.Println(state, param)
		case del:
			fmt.Println(state, param)
		case fail:
			fmt.Println(state, param)
		case quit:
			return
		}

	}

}