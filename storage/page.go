package storage

import "fmt"

const BLOCK_SIZE = 5

type Page struct {
	ID        int
	Size      int
	Content   [BLOCK_SIZE]byte
	Documents []*Document
}

func NewPage(id int) *Page {
	return &Page{
		ID: id,
	}
}

func (p *Page) AddDocument(doc *Document) error {

	totalSize := p.Size + doc.Length

	// Caso o documento caiba
	if p.Size == 0 && doc.Length == 5 {
		p.Size += doc.Length
		p.Documents = append(p.Documents, doc)
		copy(p.Content[:], doc.Content)
		return nil
	}

	if totalSize > BLOCK_SIZE {
		return fmt.Errorf("overflow na página %d, ao tentar inserir o documento de conteúdo %+v", p.ID, doc.Content)
	}

	p.Documents = append(p.Documents, doc)

	aux := 0
	for i := p.Size; i < totalSize; i++ {
		p.Content[i] = doc.Content[aux]
		aux++
	}

	p.Size += doc.Length

	return nil
}
