/*create database SGE;
go
use SGE;
go
*/
CREATE TABLE Categorias(

    ID_Categoria int PRIMARY KEY IDENTITY(1,1),
    nome VARCHAR(15) not NULL unique
);
 go

create table Produtos(

    ID_Produto int PRIMARY key IDENTITY(1,1),
    ID_Categoria int not null,
	ID_Status int not null,
    nome VARCHAR(50) not null,
    preco DECIMAL not null,
    lote VARCHAR(20) not null,
    quantidade int not null
    validade DATE not NULL

);
go

CREATE table status(
    ID_Status int PRIMARY key IDENTITY(1,1),
    nome VARCHAR(10) not null unique
);

go
alter table Produtos
add CONSTRAINT fk_categoria
FOREIGN key (ID_Categoria)
REFERENCES Categorias(ID_Categoria);
go
alter table Produtos
add constraint fk_status
foreign key (ID_Status)
references status(ID_Status);


/* atualizando tabelas de produtos, categorias e status*/

alter table Produtos
add  criacao datetime, atualizao datetime ;

go

alter table Categorias
add  criacao datetime, atualizao datetime;

go

alter table Status
add  criacao datetime, atualizao datetime;

/* parte do usuario*/


create table usuario (

    ID int primary key  identity(1,1),
    nome varchar(50) not null, 
    email varchar(100) unique not null,
    senha text not null,
    criacao datetime,
    atualizao datetime, 
);

go
create table papeis (

    id int primary key identity(1,1),
    nome varchar(50) not null,
    descricao text not null,

);

go