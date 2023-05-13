package tui

type state byte

const (
	scan state = iota
	seek
	del
	add
	quit
	fail
	pages
	menu
)

func (opt state) String() string {
	switch opt {
	case scan:
		return "Todos os documentos"
	case seek:
		return "Buscar documento"
	case del:
		return "Deletar documento"
	case add:
		return "Adicionar documento"
	case pages:
		return "Visualizar páginas"
	case menu:
		return "Menu"
	case quit:
		return "Sair"
	}

	return "Erro"
}
