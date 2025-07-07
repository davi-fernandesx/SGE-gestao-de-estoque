package services

import (
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
)

type ProdutoService struct {
	repoProduto *repository.ProdutoRepository
}



func NewProdutoService(rp *repository.ProdutoRepository) *ProdutoService {

	return &ProdutoService{

		repoProduto: rp,
	}
}

// Delete implements interfaces.ServiceCrud.
func (p *ProdutoService) Delete(id int) error {
	panic("unimplemented")
}

// Save implements interfaces.ServiceCrud.
func (p *ProdutoService) Save(entity dtos.ProdutosRequest) error {

	produto:= models.Produtos{
		Nome: entity.Nome,
		Descricao: entity.Descricao,
		Preco: entity.Preco,
		Lote: entity.Lote,
		Quantidade: entity.Quantidade,
		Validade: entity.Validade,
		CategoriaId: entity.Categoria,
		StatusId: entity.Status,
		Criacao: time.Now(),
		Atualizacao: time.Now(),
		DataChegada: entity.DataChegada,

	}

	err:= p.repoProduto.Create(produto)
	if err != nil {
		return  err
	}


	return  nil
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