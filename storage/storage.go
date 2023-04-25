package storage

import (
	"sync"
)

const MAX_PAGES = 20

var db *storage
var lock *sync.Mutex = &sync.Mutex{}

type storage struct {
	Pages [MAX_PAGES]*Page
}

func GetStorage() *storage {

	if db == nil {
		db = &storage{}

		for i := range db.Pages {
			db.Pages[i] = NewPage(i)
		}

		return db
	}

	return db
}

func (s *storage) Scan() {
	lock.Lock()
	defer lock.Unlock()
}
