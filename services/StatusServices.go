package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
)

type StatusService struct {
	repo *repository.StatusRepository
}



func NewStatusService(rs *repository.StatusRepository) interfaces.ServiceCrud[models.Status] {

	return &StatusService{
		repo: rs,
	}

}

// Delete implements interfaces.ServiceCrud.
func (s *StatusService) Delete(id int) error {
	panic("unimplemented")
}

// Save implements interfaces.ServiceCrud.
func (s *StatusService) Save(entity models.Status) error {
	
if strings.TrimSpace(entity.Nome) == "" {

	return  fmt.Errorf("campo nome nao pode ser em branco")
}
	status:= models.Status{
		Id_status: entity.Id_status, 
		Nome: entity.Nome,
		Criacao: time.Now(),
		Atualizao: time.Now(),
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
