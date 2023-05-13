package page

import "sgbd-1/src/doc"

type Page interface {
	AddDocument(*doc.Document) error
	DeleteDocument(content []byte) error

	GetDID(content []byte) (doc.DID, error)

	Size() int
	ID() int
	GetDocuments() []*doc.Document

	IsEmpty() bool
	IsFull() bool
}
