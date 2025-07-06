package controller

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
)

type CategoriaController struct {
	
	*BaseController[models.Categorias]
	categoriaService *services.CategoriaService
}

func NewCategoriaController(sc *services.CategoriaService) *CategoriaController{

	baseController:= NewBaseController[models.Categorias](sc)
	return &CategoriaController{
		BaseController: baseController,
		categoriaService: sc,
	}
}
