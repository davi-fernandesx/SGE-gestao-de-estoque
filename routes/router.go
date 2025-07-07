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
		ProdutoRepo:= repository.NewProdutoRepository(db)
	
		log.Println("INFO: services conectado com a camada de repository")
		categoriaService:= services.NewCategoriaService(categoriaRepo)
		statusService:= services.NewStatusService(statusRepo)
		produtoService:= services.NewProdutoService(ProdutoRepo)
	

		log.Println("INFO: controllers conectado com a camada de serviços")
		categoriaController:= controller.NewCategoriaController( categoriaService)
		statusController:=controller.NewStatusController(statusService)
		produtoController:= controller.NewProdutoController(produtoService)
	

		RoutesCategorias(mux, categoriaController)
		RoutesStatus(mux, statusController)
		RoutesProduto(mux, produtoController)


		

	return  mux
}