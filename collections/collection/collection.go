package collection

import "database/sql"

type Art struct {
	Name string
	Fake bool
}

type Collection struct {
	db *sql.DB
}

func NewCollection(db *sql.DB) *Collection {
	return &Collection{
		db: db,
	}
}

func (db *Collection) Add(art *Art) error {
	return nil
}

func (db *Collection) List() ([]*Art, error) {
	return nil, nil
}

func (db *Collection) Get() (*Art, error) {
	return nil, nil
}
