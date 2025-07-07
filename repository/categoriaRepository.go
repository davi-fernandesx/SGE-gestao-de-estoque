package repository

import (
	"database/sql"
	"fmt"


	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type CategoriasRepository struct {
	*Repository
}

func NewCategoriaRepository(db *sql.DB) *CategoriasRepository {

	return &CategoriasRepository{
		Repository: NewRepository(db),
	}
	
}


// Create implements interfaces.RepositoryCrud.
func (c *CategoriasRepository) Create(entity models.Categorias) error {
	
	query:= `

	insert into Categorias  (nome, criacao, atualizacao) values (@nome, @criacao, @atualizacao);
	SELECT SCOPE_IDENTITY(); 
	
	`

	stmt, err:= c.Db.Prepare(query)
	if err != nil {
		return  fmt.Errorf("erro ao preparar o statement de Create em categorias: %v", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(

		sql.Named("nome", entity.Nome),
		sql.Named("criacao", entity.Criacao),
		sql.Named("atualizacao", entity.Atualizacao),

	)	

	if err != nil {

		return  fmt.Errorf("erro no stmt de create em categorias: %v", err)
	}


	return  nil
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

