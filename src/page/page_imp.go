package page

import (
	"bytes"
	"fmt"
	"sgbd-1/src/doc"
)

const BLOCK_SIZE = 5

type PageImp struct {
	ID        int
	Size      int
	Documents []*doc.Document
}

func (p *PageImp) String() string {
	return fmt.Sprintf("Page{ID:%d, Size:%d, Documents:%v}", p.ID, p.Size, p.Documents)
}

func New(id int) Page {
	return &PageImp{
		ID:        id,
		Size:      0,
		Documents: []*doc.Document{},
	}
}

func (p *PageImp) GetID() int {
	return p.ID
}

func (p *PageImp) GetDocuments() []*doc.Document {
	return p.Documents
}

func (p *PageImp) AddDocument(doc *doc.Document) error {

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

func (p *PageImp) DeleteDocument(content []byte) error {
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

	return nil
}

func (p *PageImp) GetDID(content []byte) (doc.DID, error) {

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

func (p *PageImp) IsEmpty() bool {
	return p.Size == 0
}

func (p *PageImp) IsFull() bool {
	return p.Size == BLOCK_SIZE
}

func (p *PageImp) getDocPosition(doc *doc.Document) int {
	for i, d := range p.Documents {
		if d == doc {
			return i
		}
	}
	return -1
}

func (p *PageImp) deleteDocument(doc *doc.Document) {
	index := doc.Position
	p.Size -= doc.Length

	for i := range p.Documents {
		if p.Documents[i].Position > doc.Position {
			p.Documents[i].Position -= 1
		}
	}

	p.Documents = append(p.Documents[:index], p.Documents[index+1:]...)
}
