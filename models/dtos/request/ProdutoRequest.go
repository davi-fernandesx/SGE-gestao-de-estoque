package dtos

import "time"

type ProdutosRequest struct {
	
	
	Nome       string    		`json:"nome"`
	Descricao  string			`json:"descricao"`
	Preco      float64   		`json:"preco"` 
	Lote       string   		`json:"lote"`
	Quantidade int			`json:"quantidade"`
	Validade   time.Time 		`json:"validade"`
	Status	   int		`json:"status"`
	Categoria  int	`json:"categoria"`
	DataChegada time.Time `json:"dataChegada"`
}



func NewProdutoDTO(nome string, descricao string,preco float64, lote string, quantiadade int ,validade time.Time,status int , 
	categoria int, dataChegada time.Time) *ProdutosRequest {

	return& ProdutosRequest{
		 Nome: nome,
		 Descricao: descricao,
		 Preco: preco,
		 Lote: lote,
		 Quantidade: quantiadade,
		 Validade: validade,
		 Status: status,
		 Categoria: categoria,
		 DataChegada: dataChegada,
	}

}