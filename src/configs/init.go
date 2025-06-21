package configs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//struct da conexao para iniciar o banco de dados


func InitApplication() ( *sql.DB, error){

	log.Println("SGE-SISTEMA DE GESTAO DE ESTOQUE!")
	log.Print("desenvolvedores: Davi Fernandes e Paloma Brito" )
		//carregando os arquivos .env
		err:= godotenv.Load(".env")
		if err != nil {
	
			log.Printf("erro ao carregar o arquivo .env: %v", err)
		}else {
				log.Println("----------------------------------------------------------------")
				log.Println("Carregando informações do banco de dados:")
				log.Println("ARQUIVOS .ENV CARREGADOS")
		}
		
		db, err:= conn()//função conn vindo do arquivo que carrega a inicialização do banco de dados
		if err != nil {
			return nil, fmt.Errorf("erro ao chamar funcao Conn: %w", err)
		}

		
		if err != nil {

			log.Printf("erro ao iniciar conexão: %v", err)
		}

			
	return db, nil

}