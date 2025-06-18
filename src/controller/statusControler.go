package controller

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
)

type StatusController struct {
	*BaseController[models.Status]
}

func NewStatusController(ss interfaces.ServiceCrud[models.Status]) interfaces.ControllerApiRest[models.Status]{

	return &StatusController{
		BaseController: NewBaseController[models.Status](ss),
	}
}

var _ interfaces.ControllerApiRest[models.Status] = (*StatusController)(nil)