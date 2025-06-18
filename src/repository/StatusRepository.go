package repository

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type StatusRepository struct {
	*Repository
}

func NewStatusRepository(db *configs.Connection) interfaces.RepositoryCrud[models.Status] {

	return &StatusRepository{
		Repository: NewRepository(db),
	}
}

// Create implements interfaces.RepositoryCrud.
func (s *StatusRepository) Create(entity models.Status) error {
	panic("unimplemented")
}

// Delete implements interfaces.RepositoryCrud.
func (s *StatusRepository) Delete(entity models.Status) error {
	panic("unimplemented")
}

// FindAll implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindAll() ([]models.Status, error) {
	panic("unimplemented")
}

// FindById implements interfaces.RepositoryCrud.
func (s *StatusRepository) FindById(id int) (models.Status, error) {
	panic("unimplemented")
}

// Update implements interfaces.RepositoryCrud.
func (s *StatusRepository) Update(entity models.Status) error {
	panic("unimplemented")
}
