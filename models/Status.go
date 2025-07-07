package models

import (
	"time"

)

type Status struct {
	Id_status int  
	Nome      string      
	Criacao time.Time		
	Atualizacao time.Time		
}
