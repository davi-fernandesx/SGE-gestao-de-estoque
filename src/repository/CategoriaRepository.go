package repository

import (
	"database/sql"
	"fmt"


	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type CategoriasRepository struct {
	*Repository
}

func NewCategoriaRepository(db *sql.DB) interfaces.RepositoryCrud[models.Categorias] {

	return &CategoriasRepository{
		Repository: NewRepository(db),
	}
}


// Create implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Create(entity models.Categorias) error {
	return fmt.Errorf("teste")
}

// Delete implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Delete(entity models.Categorias) error {
	return fmt.Errorf("teste")
}

// FindAll implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) FindAll() ([]models.Categorias, error) {
	return nil, fmt.Errorf("teste")
}

// FindById implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) FindById(id int) (models.Categorias, error) {
	return  models.Categorias{},fmt.Errorf("teste");
}

// Update implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Update(entity models.Categorias) error {
	return fmt.Errorf("teste");
}

