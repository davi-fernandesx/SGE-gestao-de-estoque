# Sistema de Gestão de Estoque - Documentação

## Sumário

- [Início Rápido](#início-rápido)
- [Estrutura do projeto](#Estrutura-do-projeto)
- [Explicação das Pastas](#Explicação-das-Pastas)
- [Endpoints da Api](#endpoints-da-api)
- [Desenvolvedores](#desenvolvedores)


## Início Rápido

----
## Estrutura do projeto

```
SGE--gestao-de-estoque/
     ├── bin/
     ├── configs/
     ├── controller/
     ├── interfaces/
     ├── middleware/
     ├── models/
     ├── repository/
     ├── routes/
     ├── services/
     ├── utils/
     └── tests/
```
---
## Explicação das Pastas:

* **`bin/`**: Contém os **executáveis compilados** da aplicação. Após construir o projeto, o binário final geralmente é  armazenado aqui.
* **`configs/`**: Responsável por gerenciar todas as **configurações da aplicação**. Isso inclui carregar variáveis de ambiente (`.env`), configurar a conexão com o banco de dados e outras definições de ambiente.
* **`controller/`**: Camada que lida com as **requisições HTTP**. Os controllers recebem as requisições, validam os dados de entrada e coordenam com a camada de `services` para processar a lógica de negócio, retornando as respostas apropriadas ao cliente.
* **`interfaces/`**: Define os **contratos (interfaces)** entre as diferentes camadas do projeto. Isso promove a separação de responsabilidades e facilita a testabilidade, permitindo o uso de mocks.
* **`middleware/`**: Contém funções **intermediárias** que processam as requisições HTTP antes ou depois de chegarem aos controllers.
* **`models/`**: Define as **estruturas de dados** (structs Go) que representam as entidades da sua aplicação .
* **`repository/`**: Camada responsável pela **interação direta com o banco de dados**. Os repositórios abstraem a lógica de acesso a dados (CRUD - Criar, Ler, Atualizar, Deletar) do restante da aplicação.
* **`routes/`**: Onde as **rotas da API** são definidas e registradas. Mapeia os caminhos das URLs para os métodos específicos dos `controllers`.
* **`services/`**: Contém a **lógica de negócio principal** da aplicação. Os serviços orquestram as operações, utilizando os `repositories` para acessar dados e aplicando as regras de negócio.
* **`utils/`**: Armazena **funções auxiliares e utilitários genéricos** que podem ser reutilizados em várias partes do projeto,.

---



## Endpoints da API

A API do SGE expõe os seguintes endpoints principais:

| Método | URL                       | Descrição                                 |
| :----- | :------------------------ | :---------------------------------------- |
| `GET`  | `/api/categorias`         | Lista todas as categorias de produtos.    |
| `GET`  | `/api/categoria`          | Busca uma categoria específica pelo ID.   |
| `POST` | `/api/categorias`         | Cria uma nova categoria.                  |
| `PUT`  | `/api/categorias`         | Atualiza uma categoria existente.         |
| `DELETE`| `/api/categoria`         | Deleta uma categoria.                     |
| `GET`  | `/api/produtos`           | Lista todos os produtos no estoque.       |
| `GET`  | `/api/produtos`           | Busca um produto específico pelo ID.      |
| `POST` | `/api/produtos`           | Adiciona um novo produto ao estoque.      |
| `PUT`  | `/api/produtos`           | Atualiza as informações de um produto.    |
| `DELETE`| `/api/produto`           | Remove um produto do estoque.             |
| `GET`  | `/api/status`             | Lista todos os status de estoque.         |
| `GET`  | `/api/status`             | Busca um status de estoque pelo ID.       |
| `POST` | `/api/status`             | Cria um novo status de estoque.           |
| `PUT`  | `/api/status`             | Atualiza um status de estoque existente.  |
| `DELETE`| `/api/status`            | Deleta um status de estoque.              |

---

## Desenvolvedores

Este projeto está sendo desenvolvido por:

* **Davi Fernandes** - [GitHub](https://github.com/DaviFernandes034)
* **Paloma Brito**
