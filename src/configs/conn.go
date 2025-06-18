package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Connection struct {

	Db *sql.DB

}


func  conn() (*Connection, error) {
	// Obter variáveis de ambiente
	db_server:= os.Getenv("DB_SERVER")
	db_port:= os.Getenv("DB_PORT")
	db_database:= os.Getenv("DATABASE")
	db_user:=  os.Getenv("DB_USER")
	db_pass:= os.Getenv("SA_PASSWORD")
	


	//STRING CONEXÃO
		connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		db_user, db_pass, db_server, db_port, db_database)

	db, err:= sql.Open("sqlserver", connString)
	if err != nil {

		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err)
	}

	//verificando a conexão
	err = db.Ping()
	if err != nil {

		return nil, fmt.Errorf("erro ao verificar se a conexão ainda está ativa: %w", err)
	}

	log.Println("INFO: conexão ao banco de dados feita!!")

	//colocanco a conexão db, na struct
	return &Connection{
		Db: db,
	}, nil

}