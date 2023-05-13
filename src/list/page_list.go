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

func NewPageList() PageDirectory {
	return &PageList{
		len:     0,
		head:    nil,
		pageIds: []int{},
	}
}

func (l *PageList) IsEmpty() bool {
	return l.len == 0 && l.head == nil
}

func (l *PageList) Head() *node {
	return l.head
}

func (l *PageList) Add(p page.Page) error {
	for _, id := range l.pageIds {
		if p.ID() == id {
			return fmt.Errorf("já existe uma página com id '%d'", p.ID())
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
	l.pageIds = append(l.pageIds, p.ID())
	return nil
}

func (l *PageList) DeletePage(pageId int) (page.Page, error) {
	if l.head == nil {
		return nil, fmt.Errorf("não foi possível encontrar a página de id '%d'", pageId)
	}

	if l.head.value.ID() == pageId {
		p := l.head.value
		l.head = l.head.next
		l.pageIds = l.pageIds[1:]
		l.len--
		return p, nil
	}

	curr := l.head.next
	prev := l.head

	for curr != nil {
		if curr.value.ID() == pageId {
			prev.next = curr.next
			if curr.next == nil {
				l.pageIds = l.pageIds[:len(l.pageIds)-1]
			} else {
				currPageIdx := indexOf(l.pageIds, curr.value.ID())

				if currPageIdx < 0 {
					return nil, fmt.Errorf("não foi possível encontrar a página de id '%d'", pageId)
				}

				l.pageIds = append(l.pageIds[:currPageIdx], l.pageIds[currPageIdx+1:]...)
			}
			l.len--
			return curr.value, nil
		}
		prev = curr
		curr = curr.next
	}

	return nil, fmt.Errorf("não foi possível encontrar a página de id '%d'", pageId)
}

func indexOf(slc []int, id int) int {
	for i, v := range slc {
		if v == id {
			return i
		}
	}

	return -1
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
