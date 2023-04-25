package storage

import (
	"fmt"
)

const BLOCK_SIZE = 5

type Page struct {
	ID        int
	Size      int
	Documents []*Document
}

func (p Page) String() string {
	return fmt.Sprintf("Page{ID:%d, Size:%d, Documents:%v}", p.ID, p.Size, p.Documents)
}

func NewPage(id int) *Page {
	return &Page{
		ID: id,
	}
}

func (p *Page) AddDocument(doc *Document) error {

	totalSize := p.Size + doc.Length

	if p.Size == 0 && doc.Length == BLOCK_SIZE {
		doc.PageID = p.ID
		doc.Position = 0
		p.Documents = append(p.Documents, doc)
		p.Size += doc.Length

		return nil
	}

	if totalSize > BLOCK_SIZE {
		return fmt.Errorf("overflow na página %d, ao tentar inserir o documento de conteúdo %+v", p.ID, doc.Content)
	}

	p.Documents = append(p.Documents, doc)
	p.Size += doc.Length
	doc.Position = p.getDocPosition(doc)
	doc.PageID = p.ID

	return nil
}

func (p *Page) DeleteDocument(content []byte) error {
	if p.IsEmpty() {
		return fmt.Errorf("não existe nenhum documento na página %d", p.ID)
	}

	for index := range p.Documents {
		doc := p.Documents[index]
		if EqualContent(doc.Content, content) {
			p.removeDocument(doc)
			return nil
		}
	}

	return nil
}

func (p *Page) GetDID(content []byte) (DID, error) {

	if p.IsEmpty() {
		return DID{}, fmt.Errorf("não existe nenhum documento na página %d", p.ID)
	}

	for index := range p.Documents {
		doc := p.Documents[index]
		if EqualContent(doc.Content, content) {
			return doc.DID, nil
		}
	}

	return DID{}, fmt.Errorf("não foi encontrado nenhum documento com conteúdo %v na página %d", string(content), p.ID)
}

func (p *Page) IsEmpty() bool {
	return p.Size == 0
}

func (p *Page) IsFull() bool {
	return p.Size == BLOCK_SIZE
}

func (p *Page) getDocPosition(doc *Document) int {
	for i, d := range p.Documents {
		if d == doc {
			return i
		}
	}
	return -1
}

func (p *Page) removeDocument(doc *Document) {
	index := doc.Position
	p.Size -= doc.Length

	for i := range p.Documents {
		if p.Documents[i].Position > doc.Position {
			p.Documents[i].Position -= 1
		}
	}

	p.Documents = append(p.Documents[:index], p.Documents[index+1:]...)
}
