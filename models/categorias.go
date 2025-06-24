package models

import (
	"time"

	"github.com/google/uuid"
	
)

type Categorias struct {
	Id_categoria   uuid.UUID `json:"categoria_id"`
	Nome         string `json:"nome"`
	Criacao time.Time `json:"criacao"`
	Atualizao time.Time`json:"atualizacao"`
}
