package configs

import (
	"os"
	"testing"

	
)


func TestConnSucess(t *testing.T){

	os.Setenv("DB_SERVER", "localhost")
	os.Setenv("DB_PORT", "1433")
	os.Setenv("DATABASE", "SGE")
	db, err:= conn()
	if err != nil {

		t.Fatalf("erro. %v", err)
	}

	if db == nil {

		t.Fatalf("esperava uma *sql.DB")
	}
}




