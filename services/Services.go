package services

import (

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


