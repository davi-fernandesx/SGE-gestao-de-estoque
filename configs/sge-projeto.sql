create database SGE;
go
use SGE;
go

CREATE TABLE Categorias(

    ID_Categoria int PRIMARY KEY IDENTITY(1,1),
    nome VARCHAR(20) not NULL unique,
    criacao datetime,
    atualizacao datetime
);
 go

create table Produtos(

    ID_Produto uniqueidentifier PRIMARY key default newid(),
    nome VARCHAR(50) not null,
    descricao text not null,
    preco DECIMAL,
    lote VARCHAR(20) not null,
    quantidade int not null,
    validade DATE not NULL,
    ID_Categoria int not null,
	ID_Status int not null,
    criacao datetime ,
    atualizacao datetime

);
go

CREATE table status(
    ID_Status int PRIMARY key IDENTITY(1,1),
    nome VARCHAR(20) not null unique,
    criacao datetime ,
    atualizacao datetime
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
 select * from status;
 select * from Categorias;