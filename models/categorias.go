package models

import (
	"time"

	"github.com/google/uuid"
	
)

type Categorias struct {
	Id_categoria   uuid.UUID 
	Nome         string 
	Criacao time.Time
	Atualizao time.Time
}
