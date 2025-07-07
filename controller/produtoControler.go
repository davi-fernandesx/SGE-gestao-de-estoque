package controller

import (
	"encoding/json"
	"log"
	"net/http"

	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/utils"
)

type ProdutoController struct {
	produtoService *services.ProdutoService
}


func NewProdutoController(pc *services.ProdutoService) *ProdutoController{

	return &ProdutoController{
		produtoService: pc,
	}
}

func (pc *ProdutoController) Post() http.HandlerFunc{

	return  func(w http.ResponseWriter, r *http.Request) {

		var request dtos.ProdutosRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {

			log.Printf("ERRO: erro ao decodificar json: %v", err)
		    utils.ResponseJsonError(w, http.StatusBadRequest, "Dados invalidos")
			return
		}

		err = pc.produtoService.Save(request)
		if err != nil {
			log.Println("ERRO: falha em salvar os dados:", err)
			utils.ResponseJsonError(w, http.StatusInternalServerError, "Falha em salvar os dados")
			return
		}


		utils.ResponseJSON(w, http.StatusOK, "Produto adicionado com sucesso!")

	}
		
}