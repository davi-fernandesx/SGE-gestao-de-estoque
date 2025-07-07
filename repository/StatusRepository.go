package repository

import (
	"database/sql"
	"fmt"


	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type StatusRepository struct {
	*Repository
}

func NewStatusRepository(db *sql.DB) *StatusRepository {

	return &StatusRepository{
		Repository: NewRepository(db),
	}
}

// Create implements interfaces.RepositoryCrud.
func (s *StatusRepository) Create(entity models.Status) error {
	
	query:= `

		insert into status (nome, criacao, atualizacao) values (@nome, @criacao, @atualizacao);
		SELECT SCOPE_IDENTITY();

	`

	stmt, err:= s.Db.Prepare(query)
	if err != nil {

		return  fmt.Errorf("erro ao preparar query em Create Status: %v", err)
	}

	_, err = stmt.Exec(

		sql.Named("nome", entity.Nome),
		sql.Named("criacao", entity.Criacao),
		sql.Named("atualizacao", entity.Atualizacao),
	)

	defer stmt.Close()

	if err != nil {

		return fmt.Errorf("erro ao preparar statement: %v", err)
	}

	return  nil
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
