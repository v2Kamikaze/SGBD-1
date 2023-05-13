package storage

import (
	"fmt"
	"sgbd-1/src/doc"
	"sgbd-1/src/list"
	"sgbd-1/src/page"
)

const MAX_PAGES = 20

var db *storage

type storage struct {
	usedPages list.PageDirectory
	freePages list.PageDirectory
}

func GetStorage() Storage {
	if db == nil {
		db = &storage{}

		db.usedPages = list.NewPageList()
		db.freePages = list.NewPageList()

		for i := 0; i < MAX_PAGES; i++ {
			db.freePages.Add(page.New(i))
		}

		return db
	}

	return db
}

func (s *storage) FreePages() list.PageDirectory {
	return s.freePages
}

func (s *storage) UsedPages() list.PageDirectory {
	return s.usedPages
}

func (s *storage) ReadFree() {
	fmt.Println("Páginas livres")
	s.freePages.ReadAll()
}

func (s *storage) ReadUsed() {
	fmt.Println("Páginas usadas")
	s.usedPages.ReadAll()
}

func (s *storage) Scan() []*doc.Document {

	var docs []*doc.Document

	if s.usedPages.IsEmpty() {
		return docs
	}

	curr := s.usedPages.Head()

	for curr != nil {
		docs = append(docs, curr.Page().GetDocuments()...)
		curr = curr.Next()
	}

	return docs
}

func (s *storage) Seek(content []byte) (doc.DID, error) {

	if s.usedPages.IsEmpty() {
		return doc.DID{}, fmt.Errorf("não existe nenhum documento com o conteúdo '%s'", content)
	}

	curr := s.usedPages.Head()

	for curr != nil {
		if did, err := curr.Page().GetDID(content); err == nil {
			return did, err
		}
		curr = curr.Next()
	}

	return doc.DID{}, fmt.Errorf("não existe nenhum documento com o conteúdo '%s'", content)
}

func (s *storage) Delete(content []byte) error {
	if s.usedPages.IsEmpty() {
		return fmt.Errorf("não existe nenhuma página com conteúdo %s", content)
	}

	curr := s.usedPages.Head()

	for curr != nil {
		if err := curr.Page().DeleteDocument(content); err == nil {

			if curr.Page().IsEmpty() {
				page, err := s.usedPages.DeletePage(curr.Page().ID())

				if err != nil {
					return err
				}

				if err = s.freePages.Add(page); err != nil {
					return err
				}
			}
			return nil
		}

		curr = curr.Next()
	}

	return fmt.Errorf("não existe nenhuma página com conteúdo %s", content)

}
func (s *storage) Insert(content []byte) error {

	doc, err := doc.New(content)

	if err != nil {
		return err
	}

	// Se não existirem páginas sendo usadas, recuperar uma página das páginas livres e mover essa página para ser usada.
	if s.usedPages.IsEmpty() {
		page, err := s.freePages.DeletePage(s.freePages.Head().Page().ID())

		if err != nil {
			return err
		}

		if err = page.AddDocument(doc); err != nil {
			return err
		}

		if err = s.usedPages.Add(page); err != nil {
			return err
		}

		return nil
	}

	curr := s.usedPages.Head()

	for curr != nil {

		// Tinha espaço disponível na página.
		if err = curr.Page().AddDocument(doc); err == nil {
			return nil
		}

		// Se não existir uma próxima página, recuperar uma página das páginas livres e mover essa página para ser usada.
		if curr.Next() == nil {

			if s.freePages.IsEmpty() {
				return fmt.Errorf("não foi possível inserir um novo documento de tamanho '%d' e conteúdo '%s', não existe espaço nas páginas e não é possível alocar mais", len(content), content)
			}

			page, err := s.freePages.DeletePage(s.freePages.Head().Page().ID())

			if err != nil {
				return err
			}

			if err = page.AddDocument(doc); err != nil {
				return err
			}

			if err = s.usedPages.Add(page); err != nil {
				return err
			}

			break
		}
		curr = curr.Next()
	}

	return nil

}
