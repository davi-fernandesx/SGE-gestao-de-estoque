package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, statusCode int, data any){

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("ERRO: falha ao decodificar resposta Json: %v",err)

			http.Error(w, "erro interno", http.StatusInternalServerError)
		}


	}
}

func ResponseJsonError(w http.ResponseWriter, statusCode int, message string) {

	errorResponse:= map[string]string{"error": message}

	ResponseJSON(w, statusCode, errorResponse)
}