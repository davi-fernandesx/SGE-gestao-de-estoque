package models

import (
	"time"

	"github.com/google/uuid"
)

type Status struct {
	Id_status uuid.UUID  
	Nome      string      
	Criacao time.Time		
	Atualizao time.Time		
}
