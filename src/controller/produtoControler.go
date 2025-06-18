package controller

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type ProdutoController struct {
	*BaseController[models.Produtos]
}


func NewProdutoController(sp interfaces.ServiceCrud[models.Produtos]) interfaces.ControllerApiRest[models.Produtos]{

	return &ProdutoController{
		BaseController: NewBaseController[models.Produtos](sp),
	}
}

var _ interfaces.ControllerApiRest[models.Produtos] = (*ProdutoController)(nil)