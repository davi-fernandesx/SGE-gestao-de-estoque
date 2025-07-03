package dtos

import "time"

type ProdutosRequest struct {
	
	
	Nome       string    		`json:"nome"`
	Preco      float64   		`json:"preco"` 
	Lote       string   		`json:"lote"`
	Quantidade int			`json:"quantiadade"`
	Validade   time.Time 		`json:"validade"`
	Status	   int		`json:"status"`
	Categoria  int	`json:"categoria"`
}



func NewProdutoDTO(nome string, preco float64, lote string, quantiadade int ,validade time.Time,status int , categoria int) *ProdutosRequest {

	return& ProdutosRequest{
		 Nome: nome,
		 Preco: preco,
		 Lote: lote,
		 Quantidade: quantiadade,
		 Validade: validade,
		 Status: status,
		 Categoria: categoria,
	}

}