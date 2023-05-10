package page

import "sgbd-1/src/doc"

type Page interface {
	AddDocument(*doc.Document) error
	DeleteDocument([]byte) error

	GetDID(content []byte) (doc.DID, error)

	GetID() int
	GetDocuments() []*doc.Document

	IsEmpty() bool
	IsFull() bool
}
