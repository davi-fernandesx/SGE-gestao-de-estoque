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