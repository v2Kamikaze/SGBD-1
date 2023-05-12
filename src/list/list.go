package list

import "sgbd-1/src/page"

type List interface {
	IsEmpty() bool
	Head() *node
	Add(page.Page) error
	DeletePage(int) (page.Page, error)
	ReadAll()
}
