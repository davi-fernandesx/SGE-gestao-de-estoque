package routes

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/middleware"
)

func RoutesProduto(mux *http.ServeMux, ProdutoController *controller.ProdutoController) {


	mux.HandleFunc("/getProduto", middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					ProdutoController.Get()),
				),)

	mux.HandleFunc("/getAllProduto",middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					ProdutoController.GetAll()),
				),)

	mux.HandleFunc("/postProduto", middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					ProdutoController.Post()),
				),)

	log.Println("INFO: Rotas de Produto carregadas.")
	
}