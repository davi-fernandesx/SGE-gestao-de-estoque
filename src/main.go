package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/controller"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/repository"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/services"
	"github.com/joho/godotenv"
)

func main(){

	log.Println("SGE-SISTEMA DE GESTAO DE ESTOQUE!")
	log.Print("desenvolvedor: Davi fernandes e Paloma")
	//carregando os arquivos .env
	err:= godotenv.Load(".env.docker")
	if err != nil {

		log.Printf("erro ao carregar o arquivo .env: %v", err)
	}else {
			log.Println("----------------------------------------------------------------")
		    log.Println("Carregando informações do banco de dados:")
			log.Println("ARQUIVOS .ENV CARREGADOS")
	}

	

	var InitConnection configs.InitConnection //iniciando a inicializaçao do banco de dados

	db ,err := InitConnection.Init() //chamando a funcao para iniciar a conexão
	if err != nil {

		log.Printf("erro ao iniciar conexão: %v", err)
	}else{

		log.Println("INFO: CONEXÃO COM O BANCO DE DADOS, COMPLETA!")

		log.Println("INFO: aplicação iniciada com sucesso!!")
		
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

		log.Println("INFO: ROTAS HTTP DE CATEGORIA, STATUS E PRODUTOS CARREGADAS!")
		http.HandleFunc("/getCategoria", categoriaController.Get())
		http.HandleFunc("/getAllCategoria", categoriaController.GetAll())
		http.HandleFunc("/PostCategoria", categoriaController.Post())


		http.HandleFunc("/getStatus", StatusController.Get())
		http.HandleFunc("/getAllStatus", StatusController.GetAll())
		http.HandleFunc("/PostStatus", StatusController.Post())


		http.HandleFunc("/getProduto", ProdutoController.Get())
		http.HandleFunc("/getAllProduto", ProdutoController.GetAll())
		http.HandleFunc("/postProduto", ProdutoController.Post())


		log.Println("INFO: porta usada: :8080")
		log.Println("----------------------------------------------------------------")
		printMemoryUsage()
		
		http.ListenAndServe(":8080", nil)

		

	}
	
	
}



func printMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

   
    log.Printf("INFO: Memória total do sistema usada: %.2f MB\n", float64(m.Sys)/1024/1024)
}