package list

import "sgbd-1/src/page"

type node struct {
	value page.Page
	next  *node
}

func (n *node) Next() *node {
	return n.next
}

func (n *node) Page() page.Page {
	return n.value
}
