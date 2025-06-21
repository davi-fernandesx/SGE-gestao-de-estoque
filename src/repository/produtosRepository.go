package repository

import (
	"database/sql"
	"fmt"


	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type ProdutoRepository struct {
	*Repository
}


func NewProdutoRepository(db *sql.DB) interfaces.RepositoryCrud[models.Produtos] {

	return &ProdutoRepository{

		Repository: NewRepository(db),
	}
}

// Create implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Create(entity models.Produtos) error {
	return fmt.Errorf("teste")
}

// Delete implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Delete(entity models.Produtos) error {
	return fmt.Errorf("teste")
}

// FindAll implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) FindAll() ([]models.Produtos, error) {
	return nil, fmt.Errorf("teste")
}

// FindById implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) FindById(id int) (models.Produtos, error) {
	return models.Produtos{},fmt.Errorf("teste")
}

// Update implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Update(entity models.Produtos) error {
	return fmt.Errorf("teste")
}

