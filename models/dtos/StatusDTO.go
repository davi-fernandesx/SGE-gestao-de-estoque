package dtos

import "time"

type StatusDTO struct {
	Nome    string `json:"nome"`
	Criacao time.Time `json:"criacao"`
	Atualizao  time.Time `json:"atualizacao"`
}

func NewStatusDto(nome string) *StatusDTO {

	return &StatusDTO{
		Nome: nome,
	}
}