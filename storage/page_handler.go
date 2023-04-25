package storage

type PageHandler interface {
	AddDocument(*Document) error
	DeleteDocument([]byte) error
	IsEmpty() bool
	IsFull()
}
