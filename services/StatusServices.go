package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
)

type StatusService struct {
	repo *repository.StatusRepository
}



func NewStatusService(rs *repository.StatusRepository) *StatusService {

	return &StatusService{
		repo: rs,
	}

}

// Delete implements interfaces.ServiceCrud.
func (s *StatusService) Delete(id int) error {
	panic("unimplemented")
}

// Save implements interfaces.ServiceCrud.
func (s *StatusService) Save(entity dtos.StatusRequest) error {
	
if strings.TrimSpace(entity.Nome) == "" {

	return  fmt.Errorf("campo nome nao pode ser em branco")
}
	status:= models.Status{
		Nome: entity.Nome,
		Criacao: time.Now(),
		Atualizacao: time.Now(),
	}

	err:= s.repo.Create(status)
	if err != nil {
		return  err
	}

	return  nil
}

// SearchAll implements interfaces.ServiceCrud.
func (s *StatusService) SearchAll() ([]models.Status, error) {
	panic("unimplemented")
}

// SearchID implements interfaces.ServiceCrud.
func (s *StatusService) SearchID(id int) (models.Status, error) {
	panic("unimplemented")
}

// Update implements interfaces.ServiceCrud.
func (s *StatusService) Update(entity models.Status) {
	panic("unimplemented")
}
