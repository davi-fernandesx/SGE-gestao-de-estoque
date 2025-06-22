package routes

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/middleware"
)

func RoutesCategorias(mux *http.ServeMux, categoriaController *controller.CategoriaController) {

	mux.HandleFunc("/getCategoria",
			middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					categoriaController.Get()),	
					),)

	mux.HandleFunc("/getAllCategoria", 
			middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					categoriaController.GetAll()),
			))

	mux.HandleFunc("/PostCategoria",
			middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodPost},
				categoriaController.Post()),
			),)

	log.Println("INFO: Rotas de Categoria carregadas.")
	

}