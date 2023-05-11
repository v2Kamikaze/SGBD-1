package page

import (
	"bytes"
	"fmt"
	"sgbd-1/src/doc"
)

const BLOCK_SIZE = 5

type page struct {
	ID        int
	Size      int
	Documents []*doc.Document
}

func (p *page) String() string {
	return fmt.Sprintf("Page{ID:%d, Size:%d, Documents:%v}", p.ID, p.Size, p.Documents)
}

func New(id int) Page {
	return &page{
		ID:        id,
		Size:      0,
		Documents: []*doc.Document{},
	}
}

func (p *page) GetID() int {
	return p.ID
}

func (p *page) GetDocuments() []*doc.Document {
	return p.Documents
}

func (p *page) AddDocument(doc *doc.Document) error {

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

func (p *page) DeleteDocument(content []byte) error {
	if p.IsEmpty() {
		return fmt.Errorf("não existe nenhum documento na página %d", p.ID)
	}

	for index := range p.Documents {
		doc := p.Documents[index]
		if bytes.Equal(doc.Content, content) {
			p.deleteDocument(doc)
			return nil
		}
	}

	return fmt.Errorf("não existe nenhum documento com conteúdo %s na página %d", content, p.ID)
}

func (p *page) GetDID(content []byte) (doc.DID, error) {

	if p.IsEmpty() {
		return doc.DID{}, fmt.Errorf("não existe nenhum documento na página %d", p.ID)
	}

	for index := range p.Documents {
		doc := p.Documents[index]
		if bytes.Equal(doc.Content, content) {
			return doc.DID, nil
		}
	}

	return doc.DID{}, fmt.Errorf("não foi encontrado nenhum documento com conteúdo %v na página %d", string(content), p.ID)
}

func (p *page) IsEmpty() bool {
	return p.Size == 0
}

func (p *page) IsFull() bool {
	return p.Size == BLOCK_SIZE
}

func (p *page) getDocPosition(doc *doc.Document) int {
	for i, d := range p.Documents {
		if d == doc {
			return i
		}
	}
	return -1
}

func (p *page) deleteDocument(doc *doc.Document) {
	index := doc.Position
	p.Size -= doc.Length

	for i := range p.Documents {
		if p.Documents[i].Position > doc.Position {
			p.Documents[i].Position -= 1
		}
	}

	p.Documents = append(p.Documents[:index], p.Documents[index+1:]...)
}
