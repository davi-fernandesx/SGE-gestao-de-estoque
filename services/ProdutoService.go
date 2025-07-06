package services

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type ProdutoService struct {
	*BaseService[models.Produtos]
}



func NewProdutoService(rp interfaces.RepositoryCrud[models.Produtos]) interfaces.ServiceCrud[models.Produtos] {

	return &ProdutoService{

		BaseService: NewBaseService[models.Produtos](rp),
	}
}

// Delete implements interfaces.ServiceCrud.
func (p *ProdutoService) Delete(id int) error {
	panic("unimplemented")
}

// Save implements interfaces.ServiceCrud.
func (p *ProdutoService) Save(entity models.Produtos) error {
	panic("unimplemented")
}

// SearchAll implements interfaces.ServiceCrud.
func (p *ProdutoService) SearchAll() ([]models.Produtos, error) {
	panic("unimplemented")
}

// SearchID implements interfaces.ServiceCrud.
func (p *ProdutoService) SearchID(id int) (models.Produtos, error) {
	panic("unimplemented")
}

// Update implements interfaces.ServiceCrud.
func (p *ProdutoService) Update(entity models.Produtos) {
	panic("unimplemented")
}