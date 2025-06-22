package services

import (
	"fmt"
	"log"

	"github.com/DaviFernandes034/SGE--gestao-de-estoque/interfaces"
)

//baseService é um tipo generico, aonde nao preciso ficar digitando a mesma coisa no resto dos arquivos de Service
//"[t any]", Aceita qualquer model que eu passar
type BaseService[t any] struct {
	Repo interfaces.RepositoryCrud[t] //interface generica do repository contendo os metados CRUD

}
/* 
	construtor da struct, passo um dos models no lugar de "t any", e como argumento passo a interface contendo os cruds, 
	em "[t], passo um dos models criados"

	em "interfaces.ServiceCrud[t]", indica que a struct que eu retornar vai implentar a interface ServiceCrud[t]
	(contendo os metados em que todos os servives vao utilizar), novamente em [t],
	passo o model que eu quiser
*/
func NewBaseService[t any](repo interfaces.RepositoryCrud[t]) *BaseService[t] {

	return &BaseService[t]{
		Repo: repo,
	}

}


/*
	metados vindos do serviceCrud, contendo todos os metados que meus services irão usar

	"func (b *BaseService[t])" aqui eu posso trocar o "BaseService[t]" por qualquer service que eu criar, que nao ira ter alteração 
	"b" vai significar qualquer services que eu criar, como categoria, produto, ou status
*/
// Delete implements interfaces.ServiceCrud.
func (b *BaseService[t]) Delete(entity t) {
	log.Println("INFO: funcão Delete chamada")

	err:= b.Repo.Delete(entity)
	if err != nil {
		log.Printf("ERRO: erro encontrado em delete do pacote repository: %v", err)
		fmt.Errorf("erro em deletar registro: %v",err)

	}
}

// Save implements interfaces.ServiceCrud.
func (b *BaseService[t]) Save(entity t) error {
	log.Println("INFO: função Save chamada")
	err:=b.Repo.Create(entity)
	if err != nil {
		log.Printf("ERRO: erro encontrado em create do pacote repository:  %v", err)
		return fmt.Errorf("erro em salvar os dados: %w", err)
	}

	log.Println("INFO: informações salvas com sucesso!")
	return nil
}

// SearchAll implements interfaces.ServiceCrud.
func (b *BaseService[t]) SearchAll() ([]t , error) {
	
	log.Println("INFO: função SearchAll chamada")
	result ,err:=b.Repo.FindAll()
	if err != nil {
		log.Printf("ERRO: erro em FindAll: %v", err)
		return nil, fmt.Errorf("erro em buscar todos os registros: %w", err)
	}


	return result, nil
}

// SearchID implements interfaces.ServiceCrud.
func (b *BaseService[t]) SearchID(id int) (t, error){
	log.Println("INFO: função SearchAll chamada")
	result,err:=b.Repo.FindById(id)
		if err != nil {
			log.Printf("ERRO: erro ao FindById: %v", err)
			var zero t
			return zero, fmt.Errorf("erro em buscar um registro: %v", err)
	}

	return result, nil
}

// Update implements interfaces.ServiceCrud.
func (b *BaseService[t]) Update(entity t) {
	log.Println("INFO: função Update chamada")
	err:=b.Repo.Update(entity)
			if err != nil {
			log.Printf("ERRO: erro em update do pacote repository: %v", err)
			fmt.Errorf("erro em atualizar os dados: %v", err)
	}


}


var _ interfaces.ServiceCrud[any] = (*BaseService[any])(nil)