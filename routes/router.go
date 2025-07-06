package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
)

func SetupRouter(db *sql.DB) *http.ServeMux{

	mux:= http.NewServeMux()

		log.Println("----------------------------------------------------------------")

		log.Println("informações da aplicação: ")
		log.Println("INFO: repositorys conectado com o banco de dados")
		categoriaRepo:= repository.NewCategoriaRepository(db)
		statusRepo:= repository.NewStatusRepository(db)
	
		log.Println("INFO: services conectado com a camada de repository")
		categoriaService:= services.NewCategoriaService(categoriaRepo)
		statusService:= services.NewStatusService(statusRepo)
	

		log.Println("INFO: controllers conectado com a camada de serviços")
		categoriaController:= controller.NewCategoriaController(  categoriaService.(*services.CategoriaService))
		statusController:=controller.NewStatusController(statusService.(*services.StatusService))
	

		RoutesCategorias(mux, categoriaController)
		RoutesStatus(mux, statusController)


		

	return  mux
}