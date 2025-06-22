package routes

import (
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/middleware"
)

func RoutesStatus(mux *http.ServeMux, StatusController *controller.StatusController) {

	mux.HandleFunc("/getStatus", middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					StatusController.Get()),
				),
			)

	mux.HandleFunc("/getAllStatus",middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodGet},
					StatusController.GetAll()),
				),
			)

	mux.HandleFunc("/PostStatus", middleware.LoggerMiddleware(
				middleware.MethodMiddleware([]string{http.MethodPost},
					StatusController.Post()),
			),
		)
	log.Println("INFO: Rotas de Status carregadas.")
}