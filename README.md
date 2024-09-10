# Inventory Management System

Este é um sistema de gerenciamento de estoque desenvolvido em Golang, com o objetivo de praticar e aplicar conhecimentos de desenvolvimento Back-End. Ele permite gerenciar produtos, adicionar, atualizar e consultar informações de inventário.

## Funcionalidades

- Cadastro de produtos
- Atualização de informações de produtos
- Consulta de produtos no estoque
- Exclusão de produtos do inventário

## Tecnologias Utilizadas

- **Linguagem**: Golang
- **Framework**: Gin (para criação de APIs)
- **Banco de Dados**: PostgreSQL (ou SQLite para ambiente local)
- **Teste**: Testify (com suporte a mocks)
- **Docker**: Para execução do ambiente em containers (opcional)

## Instalação

### Pré-requisitos

- [Go 1.20+](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/) ou [SQLite](https://www.sqlite.org/download.html)
- [Git](https://git-scm.com/downloads)

### Passos para rodar o projeto localmente

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/inventory-management-system.git
   cd inventory-management-system

2. Instale as Dependências:
    ```bash
    go mod tidy

3. Rode a Aplicação:
    ```bash
    go run main.go

4. Estrutura do Projeto:
inventory-management-system
├── controllers/    # Contém a lógica dos endpoints
├── models/         # Modelos de dados (structs)
├── repositories/   # Lógica de acesso ao banco de dados
├── services/       # Camada de regras de negócio
├── tests/          # Arquivos de teste
├── main.go         # Arquivo principal da aplicação
├── go.mod          # Gerenciamento de dependências
├── README.md       # Documentação do projeto
└── scripts/        # Scripts para migrações e configuração inicial
