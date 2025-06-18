package repository


import (
	

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
)

type Repository struct {
	Db *configs.Connection
}

func NewRepository(db *configs.Connection)*Repository{

	
	return &Repository{
		Db: db,
	}


}