package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/utils"
)

type BaseController[t any] struct {
	service interfaces.ServiceCrud[t]
}

func NewBaseController[t any](service interfaces.ServiceCrud[t]) *BaseController[t] {

	return &BaseController[t]{
		service: service,
	}
}

func (bc *BaseController[t]) GetAll() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		result, err := bc.service.SearchAll()
		if err != nil {
			log.Printf("ERRO: erro encontrado em SearchAll do pacote service: %v", err)
			utils.ResponseJsonError(w, http.StatusBadGateway, "Erro ao buscar registros")
			return
		}

		response := map[string]any{
			"registros": result,
		}

		utils.ResponseJSON(w, http.StatusOK, response)
	}
}

func (bc *BaseController[t]) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		param := r.URL.Query().Get("id")
		if param == "" {
			log.Println("ERRO: cliente não passou o parametro id")
			utils.ResponseJsonError(w, http.StatusBadRequest, "passar o Id")
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {

			utils.ResponseJsonError(w, http.StatusInternalServerError, "Erro ao converter parametros Id")
			return
		}

		result, err := bc.service.SearchID(id)
		if err != nil {
			log.Printf("ERRO: erro encontrado em SearchID do pacote service: %v", err)
			utils.ResponseJsonError(w, http.StatusBadGateway, "Erro ao buscar o ID")
			return
		}

		response := map[string]any{
			"registro": result,
		}

		utils.ResponseJSON(w, http.StatusOK, response)
	}
}

// Delete implements interfaces.ControllerApiRest.
func (bc *BaseController[t]) Delete() http.HandlerFunc {
	panic("unimplemented")
}

// Update implements interfaces.ControllerApiRest.
func (bc *BaseController[t]) Update() http.HandlerFunc {
	panic("unimplemented")
}
