package routes

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/middleware"
)

func RoutesProduto(mux *http.ServeMux, ProdutoController *controller.ProdutoController) {


		mux.HandleFunc("/api/produto",
			middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodPost},
				ProdutoController.Post()),))

				
	log.Println("INFO: Rotas de Produto carregadas.")
	
}