package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
)


type CategoriaService struct {

	repoCategoria *repository.CategoriasRepository
}


func NewCategoriaService(repo *repository.CategoriasRepository) *CategoriaService {

	return &CategoriaService{
		repoCategoria: repo,
	}
}



func (cs *CategoriaService) Save(entityDtoRequest dtos.CategoriasRequest) error {

	if strings.TrimSpace(entityDtoRequest.Nome) == "" {

		return  fmt.Errorf("nome da categoria não pode estar vazio")
	}

	categoria:= models.Categorias{
		Nome: entityDtoRequest.Nome,
		Criacao: time.Now(),
		Atualizacao: time.Now(),
	}


	err:=cs.repoCategoria.Create(categoria)
	if err != nil {

		return  err
	}

	return nil
}


func (cs *CategoriaService) SearchID(id int) (models.Categorias, error){

	return models.Categorias{}, nil
}

func (cs *CategoriaService) Delete(id int)  (error){

	return  nil
}

func (cs *CategoriaService) SearchAll() ([]models.Categorias, error){

	return nil, nil
}

func (cs *CategoriaService) Update(entity dtos.CategoriasRequest){}
