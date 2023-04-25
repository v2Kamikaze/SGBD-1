package storage

type StorageHandler interface {
	Scan() []*Document
	Seek([]byte)
	Delete([]byte)
	Insert([]byte)
}
