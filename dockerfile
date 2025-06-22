##build
FROM golang:1.22.2-alpine AS builder

#instalando dependencias
RUN apk add --no-cache git

WORKDIR /app 

##arquivos de dependencia
COPY go.mod go.sum ./
RUN go mod download

##copiando o resto do codigo e copilando
COPY . .
RUN  go build -ldflags="-s -w" -o meuApp

#garantindo o binario executavel
RUN chmod +x /app/meuApp    

#port
EXPOSE 8080
# Instala dependências para SQL Server (ODBC)
RUN apk add --no-cache freetds unixodbc ca-certificates

#execução
CMD [ "/app/meuApp"]

