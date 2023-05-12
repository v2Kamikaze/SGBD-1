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
	UsedPages list.List
	FreePages list.List
}

func GetStorage() Storage {
	if db == nil {
		db = &storage{}

		db.UsedPages = list.NewPageList()
		db.FreePages = list.NewPageList()

		for i := 0; i < MAX_PAGES; i++ {
			db.FreePages.Add(page.New(i))
		}

		return db
	}

	return db
}

func (s *storage) ReadFree() {
	fmt.Println("Páginas livres")
	s.FreePages.ReadAll()
}

func (s *storage) ReadUsed() {
	fmt.Println("Páginas usadas")
	s.UsedPages.ReadAll()
}

func (s *storage) Scan() []*doc.Document {

	var docs []*doc.Document

	if s.UsedPages.IsEmpty() {
		return docs
	}

	curr := s.UsedPages.Head()

	for curr != nil {
		docs = append(docs, curr.Page().GetDocuments()...)
		curr = curr.Next()
	}

	return docs
}

func (s *storage) Seek(content []byte) (doc.DID, error) {

	if s.UsedPages.IsEmpty() {
		return doc.DID{}, fmt.Errorf("não existe nenhum documento com o conteúdo '%s'", content)
	}

	curr := s.UsedPages.Head()

	for curr != nil {
		if did, err := curr.Page().GetDID(content); err == nil {
			return did, err
		}
		curr = curr.Next()
	}

	return doc.DID{}, fmt.Errorf("não existe nenhum documento com o conteúdo '%s'", content)
}

func (s *storage) Delete(content []byte) error {
	if s.UsedPages.IsEmpty() {
		return fmt.Errorf("não existe nenhuma página com conteúdo %s", content)
	}

	curr := s.UsedPages.Head()

	for curr != nil {
		if err := curr.Page().DeleteDocument(content); err == nil {

			if curr.Page().IsEmpty() {
				page, err := s.UsedPages.DeletePage(curr.Page().GetID())

				if err != nil {
					return err
				}

				if err = s.FreePages.Add(page); err != nil {
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
	if s.UsedPages.IsEmpty() {
		page, err := s.FreePages.DeletePage(s.FreePages.Head().Page().GetID())

		if err != nil {
			return err
		}

		if err = page.AddDocument(doc); err != nil {
			return err
		}

		if err = s.UsedPages.Add(page); err != nil {
			return err
		}

		return nil
	}

	curr := s.UsedPages.Head()

	for curr != nil {

		// Tinha espaço disponível na página.
		if err = curr.Page().AddDocument(doc); err == nil {
			return nil
		}

		// Se não existir uma próxima página, recuperar uma página das páginas livres e mover essa página para ser usada.
		if curr.Next() == nil {

			if s.FreePages.IsEmpty() {
				return fmt.Errorf("não foi possível inserir um novo documento de tamanho '%d' e conteúdo '%s', não existe espaço nas páginas e não é possível alocar mais", len(content), content)
			}

			page, err := s.FreePages.DeletePage(s.FreePages.Head().Page().GetID())

			if err != nil {
				return err
			}

			if err = page.AddDocument(doc); err != nil {
				return err
			}

			if err = s.UsedPages.Add(page); err != nil {
				return err
			}

			break
		}
		curr = curr.Next()
	}

	return nil

}
