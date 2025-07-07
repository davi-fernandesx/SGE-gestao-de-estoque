package controller

import (
	"encoding/json"
	"log"
	"net/http"

	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/utils"
)

type CategoriaController struct {

	categoriaService *services.CategoriaService
}

func NewCategoriaController(sc *services.CategoriaService) *CategoriaController{

	return &CategoriaController{

		categoriaService: sc,
	}
}

func (cc *CategoriaController) Post() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var request dtos.CategoriasRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {

			log.Printf("ERRO: erro ao decodificar json: %v", err)
		    utils.ResponseJsonError(w, http.StatusBadRequest, "Dados invalidos")
			return
		}

		err = cc.categoriaService.Save(request)
		if err != nil {
			log.Println("ERRO: falha em salvar os dados:", err)
			utils.ResponseJsonError(w, http.StatusInternalServerError, "Falha em salvar os dados")
			return
		}


		utils.ResponseJSON(w, http.StatusOK, "Categoria adicionada com sucesso!")

	}

}