package response

import (
	"time"

	dtos "github.com/DaviFernandes034/SGE--gestao-de-estoque/models/dtos/request"
)

type ProdutoResponse struct {
	Id       int `json:"id"`
	Nome     string `json:"nome"`
	Preco    int `json:"preco"`
	Lote     string `json:"lote"`
	Validade time.Time `json:"validade"`
	Quantidade int `json:"quantidade"`
	Categoria dtos.CategoriasRequest `json:"categoria"`
	Status dtos.StatusRequest `json:"status"`

}
