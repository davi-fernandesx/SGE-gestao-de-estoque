package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
)


type CategoriaService struct {

	repoCategoria *repository.CategoriasRepository
}


func NewCategoriaService(repo *repository.CategoriasRepository) interfaces.ServiceCrud[models.Categorias] {

	return &CategoriaService{
		repoCategoria: repo,
	}
}

func (cs *CategoriaService) Save(entity models.Categorias) error {

	if strings.TrimSpace(entity.Nome) == "" {

		return  fmt.Errorf("nome da categoria não pode estar vazio")
	}

	categoria:= models.Categorias{
		Id_categoria: entity.Id_categoria,
		Nome: entity.Nome,
		Criacao: time.Now(),
		Atualizao: time.Now(),
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

func (cs *CategoriaService) Update(entity models.Categorias){}


var _ interfaces.ServiceCrud[models.Categorias] = (*CategoriaService)(nil)