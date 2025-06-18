package controller

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type CategoriaController struct {
	*BaseController[models.Categorias]
}

func NewCategoriaController(sc interfaces.ServiceCrud[models.Categorias]) interfaces.ControllerApiRest[models.Categorias]{

	return &CategoriaController{
		BaseController: NewBaseController[models.Categorias](sc),
	}
}

var _ interfaces.ControllerApiRest[models.Categorias] = (*CategoriaController)(nil)