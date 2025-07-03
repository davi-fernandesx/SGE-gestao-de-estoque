package models

import (
	"time"

	"github.com/google/uuid"
)

type Produtos struct{

	Id_produto uuid.UUID 
	Nome string 
	Preco float64 
	Lote string 
	Quantidade int
    Validade time.Time 
	CategoriaId int64 
	StatusId int64 
	Criacao time.Time
	Atualizao time.Time

}


