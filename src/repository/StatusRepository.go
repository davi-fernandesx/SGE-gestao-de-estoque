package repository

import (
	"database/sql"
	"fmt"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type StatusRepository struct {
	*Repository
}

func NewStatusRepository(db *sql.DB) interfaces.RepositoryCrud[models.Status] {

	return &StatusRepository{
		Repository: NewRepository(db),
	}
}

// Create implements interfaces.RepositoryCrud.
func (s *StatusRepository) Create(entity models.Status) error {
	return fmt.Errorf("teste")
}

// Delete implements interfaces.RepositoryCrud.
func (s *StatusRepository) Delete(entity models.Status) error {
	return fmt.Errorf("teste")
}

// FindAll implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindAll() ([]models.Status, error) {
	return []models.Status{},fmt.Errorf("teste")
}

// FindById implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindById(id int) (models.Status, error) {
	return models.Status{}, fmt.Errorf("teste")
}

// Update implements interfaces.RepositoryCrud.
func (s *StatusRepository) Update(entity models.Status) error {
	return fmt.Errorf("teste")
}
