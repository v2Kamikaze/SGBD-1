package storage

import "sgbd-1/src/doc"

type StorageHandler interface {
	Scan() []*doc.Document
	Seek([]byte)
	Delete([]byte)
	Insert([]byte)
}
