package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)



type BaseController[t any] struct {
	service interfaces.ServiceCrud[t]
}


func NewBaseController[t any](service interfaces.ServiceCrud[t]) *BaseController[t] {


	return &BaseController[t]{
		service: service,
	}
}

func (bc *BaseController[t]) GetAll() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("INFO: função GetAll do controller chamada")
		if r.Method != http.MethodGet {

			http.Error(w, "request não aceito, request Get esperado", http.StatusMethodNotAllowed)
			return
		}

		result, err := bc.service.SearchAll()
		if err != nil {
			log.Printf("ERRO: erro encontrado em SearchAll do pacote service: %w", err)
			http.Error(w, "erro ao buscar os registros", http.StatusBadGateway)
			return
		}

		response := map[string]any{
			"registros": result,
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (bc *BaseController[t]) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("INFO: função Get do controller chamada")
		if r.Method != http.MethodGet {

			http.Error(w, "request nao aceito, request Get esperado", http.StatusMethodNotAllowed)
			return

		}

		param := r.URL.Query().Get("id")
		if param == "" {
			log.Println("ERRO: cliente não passou o parametro id")
			http.Error(w, "favor, passar um id", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(param)
		if err != nil {

			http.Error(w, "erro ao transformar o parametro id em tipo inteiro", http.StatusInternalServerError)
			return
		}

		result, err := bc.service.SearchID(id)
		if err != nil {
			log.Printf("ERRO: erro encontrado em SearchID do pacote service: %w", err)
			http.Error(w, "erro ao buscar registro", http.StatusBadGateway)
		}

		response := map[string]any{
			"registro": result,
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (bc *BaseController[t]) Post() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("INFO: função Post do pacote controller chamada")
		if r.Method != http.MethodPost {

			http.Error(w, "request nao aceito, request Post esperado", http.StatusMethodNotAllowed)
			return

		}
		var request t

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {

			log.Printf("ERRO: erro ao decodificar json: %w", err)
			http.Error(w, "dados invalidos", http.StatusBadRequest)
			return
		}

		err = bc.service.Save(request)
		if err != nil {
			log.Println("ERRO: falha em salvar os dados")
			http.Error(w, "erro interno ao processar requisição", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(request)

	}

}

// Delete implements interfaces.ControllerApiRest.
func (bc *BaseController[t]) Delete() http.HandlerFunc {
	panic("unimplemented")
}

// Update implements interfaces.ControllerApiRest.
func (bc *BaseController[t]) Update() http.HandlerFunc {
	panic("unimplemented")
}

var _ interfaces.ControllerApiRest[any] = (*BaseController[any])(nil)

