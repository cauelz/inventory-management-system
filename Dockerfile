# Usar uma imagem base do Go
FROM golang:1.22-alpine

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte da aplicação
COPY . .

# Compilar a aplicação
RUN go build -o main .

# Definir a porta que a aplicação irá expor
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]