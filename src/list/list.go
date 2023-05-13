package list

import "sgbd-1/src/page"

type PageDirectory interface {
	IsEmpty() bool
	Head() *node
	Add(page.Page) error
	DeletePage(int) (page.Page, error)
	ReadAll()
}
