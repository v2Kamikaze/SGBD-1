package storage

import (
	"sgbd-1/src/doc"
	"sgbd-1/src/list"
)

type Storage interface {
	Scan() []*doc.Document
	Seek(content []byte) (doc.DID, error)
	Delete(content []byte) error
	Insert(content []byte) error

	FreePages() list.PageDirectory
	UsedPages() list.PageDirectory

	ReadFree()
	ReadUsed()
}
