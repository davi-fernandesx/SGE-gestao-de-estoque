package controller

import (
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/models"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
)

type StatusController struct {
	*BaseController[models.Status]
	statusService *services.StatusService
}

func NewStatusController(ss *services.StatusService) *StatusController{

	baseController:= NewBaseController[models.Status](ss)
	return &StatusController{
		BaseController: baseController,
		statusService: ss,
	}
}

var _ interfaces.ControllerApiRest[models.Status] = (*StatusController)(nil)