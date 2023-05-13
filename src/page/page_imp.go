package page

import (
	"bytes"
	"fmt"
	"sgbd-1/src/doc"
)

const BLOCK_SIZE = 5

type page struct {
	id        int
	size      int
	documents []*doc.Document
}

func (p *page) String() string {
	return fmt.Sprintf("Page{ID:%d, size:%d, Documents:%v}", p.id, p.size, p.documents)
}

func New(id int) Page {
	return &page{
		id:        id,
		size:      0,
		documents: []*doc.Document{},
	}
}

func (p *page) ID() int {
	return p.id
}

func (p *page) Size() int {
	return p.size
}

func (p *page) GetDocuments() []*doc.Document {
	return p.documents
}

func (p *page) AddDocument(doc *doc.Document) error {

	totalsize := p.size + doc.Length

	if p.size == 0 && doc.Length == BLOCK_SIZE {
		doc.PageID = p.id
		doc.Position = 0
		p.documents = append(p.documents, doc)
		p.size += doc.Length

		return nil
	}

	if totalsize > BLOCK_SIZE {
		return fmt.Errorf("overflow na página %d, ao tentar inserir o documento de conteúdo %+v", p.id, doc.Content)
	}

	p.documents = append(p.documents, doc)
	p.size += doc.Length
	doc.Position = p.getDocPosition(doc)
	doc.PageID = p.id

	return nil
}

func (p *page) DeleteDocument(content []byte) error {
	if p.IsEmpty() {
		return fmt.Errorf("não existe nenhum documento na página %d", p.id)
	}

	for index := range p.documents {
		doc := p.documents[index]
		if bytes.Equal(doc.Content, content) {
			p.deleteDocument(doc)
			return nil
		}
	}

	return fmt.Errorf("não existe nenhum documento com conteúdo %s na página %d", content, p.id)
}

func (p *page) GetDID(content []byte) (doc.DID, error) {

	if p.IsEmpty() {
		return doc.DID{}, fmt.Errorf("não existe nenhum documento na página %d", p.id)
	}

	for index := range p.documents {
		doc := p.documents[index]
		if bytes.Equal(doc.Content, content) {
			return doc.DID, nil
		}
	}

	return doc.DID{}, fmt.Errorf("não foi encontrado nenhum documento com conteúdo %v na página %d", string(content), p.id)
}

func (p *page) IsEmpty() bool {
	return p.size == 0
}

func (p *page) IsFull() bool {
	return p.size == BLOCK_SIZE
}

func (p *page) getDocPosition(doc *doc.Document) int {
	for i, d := range p.documents {
		if d == doc {
			return i
		}
	}
	return -1
}

func (p *page) deleteDocument(doc *doc.Document) {
	index := doc.Position
	p.size -= doc.Length

	for i := range p.documents {
		if p.documents[i].Position > doc.Position {
			p.documents[i].Position -= 1
		}
	}

	p.documents = append(p.documents[:index], p.documents[index+1:]...)
}
