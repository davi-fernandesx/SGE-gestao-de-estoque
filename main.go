package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/configs"
	"github.com/DaviFernandes034/SGE--gestao-de-estoque/routes"
)

	func main(){

			
		db, err:= configs.InitApplication()
		if err != nil {
			log.Fatal(err)
		}
			mux:= routes.SetupRouter(db)

			log.Println("INFO: porta usada: :8080")
			log.Println("INFO: aplicação iniciada com sucesso!!")
			log.Println("----------------------------------------------------------------")
			printMemoryUsage()
			
			log.Fatal(http.ListenAndServe(":8080", mux))

	
	}


func printMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

   
    log.Printf("INFO: Memória total do sistema usada: %.2f MB\n", float64(m.Sys)/1024/1024)
}



/*
		4. Adicione Funcionalidades Essenciais (Além do Básico)
Mesmo para um MVP (Produto Mínimo Viável), algumas funcionalidades são quase obrigatórias em sistemas de estoque para serem vendáveis:

Autenticação e Autorização: Controle de acesso por usuário e papel (administrador, operador, etc.). Essencial para segurança e colaboração.
Relatórios Básicos: Listas de produtos, movimentações, estoque atual, produtos com estoque baixo. Não precisam ser complexos, mas precisam existir.
Busca e Filtros Eficientes: Facilitar a localização de produtos.
Histórico de Movimentações: Registrar quem fez o quê, quando e onde.
Contagem de Estoque: Ferramenta para facilitar a contagem física e ajuste do estoque
*/