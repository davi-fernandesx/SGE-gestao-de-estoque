package repository

import (
	"database/sql"
	"fmt"


	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type ProdutoRepository struct {
	*Repository
}


func NewProdutoRepository(db *sql.DB) *ProdutoRepository {

	return &ProdutoRepository{

		Repository: NewRepository(db),
	}
}

// Create implements interfaces.RepositoryCrud.
func (p *ProdutoRepository) Create(entity models.Produtos) error {
	
	query:= `

		insert into Produtos (nome, descricao, preco, lote, quantidade, validade, ID_Categoria, ID_Status, criacao, atualizacao, dataChegada)
		values ( @nome, @descricao,@preco, @lote,@quantidade, @validade, @id_categoria, @id_status, @criacao, @atualizacao, @dataChegada);
		 SELECT SCOPE_IDENTITY(); 
	`

	stmt, err:= p.Db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		sql.Named("nome", entity.Nome),
		sql.Named("descricao", entity.Descricao),
		sql.Named("preco", entity.Preco),
		sql.Named("lote", entity.Lote),
		sql.Named("quantidade", entity.Quantidade),
		sql.Named("validade", entity.Validade),
		sql.Named("id_categoria", entity.CategoriaId),
		sql.Named("id_status", entity.StatusId),
		sql.Named("criacao", entity.Criacao),
		sql.Named("atualizacao", entity.Atualizacao),
		sql.Named("dataChegada", entity.DataChegada),

	)

	defer stmt.Close()

	if err != nil{
		return  err
	}

	return  nil

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

