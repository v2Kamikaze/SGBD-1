package storage

import (
	"fmt"
)

type DID struct {
	PageID   int
	Position int
	Length   int
}

func (d DID) String() string {
	return fmt.Sprintf("DID{PageID:%d, Position:%d, Length:%d}", d.PageID, d.Position, d.Length)
}

type Document struct {
	DID
	Content []byte
}

func (doc Document) String() string {
	return fmt.Sprintf("Document{DID:%s, Content:%s}", doc.DID, doc.Content)
}

func NewDocument(content []byte) (*Document, error) {
	len := len(content)
	if len < 1 || len > 5 {
		return nil, fmt.Errorf("tamanho do documento fora dos limites. tamanho: %d", len)
	}

	did := DID{Length: len}

	return &Document{did, content}, nil
}
