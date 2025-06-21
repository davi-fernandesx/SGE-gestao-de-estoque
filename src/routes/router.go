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
			log.Println("INFO: repository conectado com o banco de dados")
			categoriaRepo:= repository.NewCategoriaRepository(db)
			StatusRepo:= repository.NewStatusRepository(db)
			ProdutoRepo:= repository.NewProdutoRepository(db)

			log.Println("INFO: services conectado com a camada de repository")
			categoriaService:= services.NewCategoriaService(categoriaRepo)
			StatusService:= services.NewStatusService(StatusRepo)
			ProdutoService:= services.NewProdutoService(ProdutoRepo)

			log.Println("INFO: controller conectado com a camada de serviços")
			categoriaController:= controller.NewCategoriaController(categoriaService)
			StatusController:= controller.NewStatusController(StatusService)
			ProdutoController:= controller.NewProdutoController(ProdutoService)


			RoutesCategorias(mux, categoriaController)
			RoutesProduto(mux, ProdutoController)
			RoutesStatus(mux,StatusController)

		

	return  mux
}