package mysql

import (
	"databae/sql"

	"github.com/santiaguito-g/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, created, expires string) (int, err) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
