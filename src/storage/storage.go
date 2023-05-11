package storage

import "sgbd-1/src/doc"

type Storage interface {
	Scan() []*doc.Document
	Seek(content []byte) (doc.DID, error)
	Delete(content []byte) error
	Insert(content []byte) error

	ReadFree()
	ReadUsed()
}
