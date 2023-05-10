package storage

import (
	"sgbd-1/src/list"
	"sync"
)

const MAX_PAGES = 20

var db *storage
var lock *sync.Mutex = &sync.Mutex{}

type storage struct {
	Pages     *list.PageList
	FreePages *list.PageList
}

func GetStorage() *storage {
	if db == nil {
		db = &storage{}

		db.FreePages = list.NewPageList()
		db.FreePages = list.NewPageList()

		return db
	}

	return db
}

func (s *storage) Scan() {
	lock.Lock()
	defer lock.Unlock()
}
