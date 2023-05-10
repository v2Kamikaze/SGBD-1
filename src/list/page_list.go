package list

import (
	"fmt"
	"sgbd-1/src/page"
)

type PageList struct {
	len     int
	head    *node
	pageIds []int
}

type node struct {
	value page.Page
	next  *node
}

func NewPageList() *PageList {
	return &PageList{
		len:     0,
		head:    nil,
		pageIds: []int{},
	}
}

func (l *PageList) Add(p page.Page) error {
	for _, id := range l.pageIds {
		if p.GetID() == id {
			return fmt.Errorf("já existe uma página com id '%d'", p.GetID())
		}
	}

	newNode := &node{
		value: p,
		next:  nil,
	}

	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}

	l.len++
	l.pageIds = append(l.pageIds, p.GetID())
	return nil
}

func (l *PageList) DeletePage(pageId int) (page.Page, error) {
	if l.head == nil {
		return nil, fmt.Errorf("não foi possível encontrar a página de id '%d'", pageId)
	}

	if l.head.value.GetID() == pageId {
		p := l.head.value
		l.head = l.head.next
		l.pageIds = l.pageIds[1:]
		l.len--
		return p, nil
	}

	curr := l.head.next
	prev := l.head

	for curr != nil {
		if curr.value.GetID() == pageId {
			prev.next = curr.next
			if curr.next == nil {
				l.pageIds = l.pageIds[:len(l.pageIds)-1]
			} else {
				l.pageIds = append(l.pageIds[:curr.value.GetID()-1], l.pageIds[curr.value.GetID():]...)
			}
			l.len--
			return curr.value, nil
		}
		prev = curr
		curr = curr.next
	}

	return nil, fmt.Errorf("não foi possível encontrar a página de id '%d'", pageId)
}

func (l *PageList) ReadAll() {
	fmt.Printf("PageIDS: %+v\n", l.pageIds)
	fmt.Printf("Length: %+v\n", l.len)
	fmt.Printf("Head: %+v\n", l.head)

	curr := l.head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}