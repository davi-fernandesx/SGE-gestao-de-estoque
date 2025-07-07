package routes

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/middleware"
)

func RoutesCategorias(mux *http.ServeMux, categoriaController *controller.CategoriaController) {



	mux.HandleFunc("/api/categoria",
			middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodPost},
				categoriaController.Post()),
			),)

	log.Println("INFO: Rotas de Categoria carregadas.")
	

}