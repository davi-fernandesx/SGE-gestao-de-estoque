package models

import (
	"time"

	"github.com/google/uuid"
)

type Status struct {
	Id_status uuid.UUID  `json:"status_id"`
	Nome      string      `json:"Nome"`
	Criacao time.Time		`json:"criacao"`
	Atualizao time.Time		`json:"atualizacao"`
}
