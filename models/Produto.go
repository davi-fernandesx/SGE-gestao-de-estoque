package models

import (
	"time"

	"github.com/google/uuid"
)

type Produtos struct{

	Id_produto uuid.UUID 
	Nome string 
	Descricao string
	Preco float64 
	Lote string 
	Quantidade int
    Validade time.Time 
	CategoriaId int
	StatusId int
	Criacao time.Time
	Atualizacao time.Time
	DataChegada time.Time

}


