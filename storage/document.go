package storage

import (
	"fmt"
)

type DID struct {
	PageID   int
	Position int
	Length   int
}

type Document struct {
	DID
	Content []byte
}

func NewDocument(content []byte) (*Document, error) {
	len := len(content)
	if len < 1 || len > 5 {
		return nil, fmt.Errorf("tamanho do documento fora dos limites. tamanho: %d", len)
	}

	did := DID{Length: len}

	return &Document{did, content}, nil
}
