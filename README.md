# Documentação da API Go-API

## 📖 Descrição
Este projeto é uma API RESTful desenvolvida em **Go**, projetada utilizando **Arquitetura em Camadas (Clean Architecture)**. A aplicação oferece um CRUD completo para o gerenciamento de **informações pessoais/perfil**.

---

## 🛠️ Linguagem e Tecnologias Utilizadas

- **Linguagem:** [Go (Golang)](https://go.dev/) (v1.22+)
- **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) — Framework HTTP de alta performance para Go.
- **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/) (Hospedado na plataforma Supabase).
- **Driver de Banco de Dados:** `github.com/lib/pq` (Consultas SQL nativas através da biblioteca standard `database/sql`).
- **Containerização:** Docker e Docker Compose.

---

## 🏗️ Arquitetura da Aplicação

A aplicação adota o padrão de **Arquitetura em Camadas (Clean Architecture)**, promovendo baixo acoplamento e separação clara de responsabilidades:

```
┌─────────────────────────────────────────────────────────┐
│                 cmd/main.go (Roteamento)                │
└───────────────────────────┬─────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                     Camada Controller                   │
│   (Recebe requisições HTTP e valida dados de entrada)   │
└───────────────────────────┬─────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                      Camada UseCase                     │
│           (Contém as regras de negócio)                 │
└───────────────────────────┬─────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│                    Camada Repository                    │
│    (Executa consultas SQL nativas no banco de dados)    │
└───────────────────────────┬─────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────┐
│               Banco de Dados (PostgreSQL)               │
└───────────────────────────┬─────────────────────────────┘
```

### Componentes Principais:
1. **Controller (`controller/`)**: Intercala a comunicação HTTP com os Casos de Uso. Formata as respostas JSON e trata códigos de status HTTP (200, 201, 400, 404, 500).
2. **UseCase (`usecase/`)**: Implementa a lógica de negócios e as regras da aplicação.
3. **Repository (`repository/`)**: Abstrai o acesso aos dados e interage diretamente com o banco PostgreSQL por meio de queries SQL (`SELECT`, `INSERT`).
4. **Model (`model/`)**: Define as estruturas de dados (structs) mapeadas em JSON.
5. **Mocks (`mocks/`)**: Contém os mocks das interfaces para utilização nos testes unitários.
5. **DB (`db/`)**: Gerencia e estabelece a conexão com a base de dados PostgreSQL.

---

## 🗄️ Banco de Dados

A API conecta-se a um banco de dados **PostgreSQL** utilizando consultas SQL nativas (`database/sql` + driver `github.com/lib/pq`).

### Estrutura das Tabelas:

#### Tabela `personal`
- `id` (INTEGER / SERIAL - Chave primária)
- `address`, `city`, `neighborhood`, `state`, `cep`, `phone`, `email`, `website`, `linkedin`, `github` (VARCHAR)

---

## 📂 Estrutura do Projeto

```text
go-api/
├── cmd/
│   └── main.go                     # Ponto de entrada e registro de rotas
├── controller/
│   ├── personal_controller.go      # Controller para a entidade Personal
│   └── personal_controller_test.go # Testes do controller
├── db/
│   └── conn.go                     # Conexão com o banco de dados
├── mocks/
│   ├── personal_repository_mock.go # Mock do repositório Personal
│   └── personal_usecase_mock.go    # Mock do use case Personal
├── model/
│   └── personal.go                 # Struct do modelo Personal
├── repository/
│   ├── personal_repository.go           # Implementação do repositório
│   ├── personal_repository_interface.go # Interface do repositório
│   └── personal_repository_test.go    # Testes do repositório
├── usecase/
│   ├── personal_usecase.go          # Regras de negócio da entidade Personal
│   └── personal_usecase_test.go     # Testes do use case
├── Dockerfile                  # Configuração para geração de imagem Docker
├── docker-compose.yml          # Orquestração do serviço da API
├── go.mod                      # Dependências do módulo Go
├── go.sum                      # Checksum das dependências
└── README.md                   # Documentação completa do projeto
```

---

## 🚀 Como Executar o Projeto

### Pré-requisitos
- **Go 1.22+** (ou Docker/Docker Compose instalado)
- Conexão de rede ativa para acessar o banco PostgreSQL no Supabase

### Opção 1: Executar Localmente com Go

1. **Clonar o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/go-api.git
   cd go-api
   ```

2. **Baixar as dependências:**
   ```bash
   go mod download
   ```

3. **Iniciar a aplicação:**
   ```bash
   go run cmd/main.go
   ```
   *A API estará acessível na porta 8000: `http://localhost:8000`*

---

### Opção 2: Executar via Docker Compose

1. **Build e inicialização do container:**
   ```bash
   docker-compose up --build
   ```
   *A API estará acessível em `http://localhost:8000`*

---

## 📡 Endpoints da API

### 🟢 Health Check

#### `GET /ping`
Endpoint de verificação de status da API.
- **Resposta Sucesso (`200 OK`):**
  ```json
  {
    "message": "pong"
  }
  ```

---

###  Informações Pessoais (`/personal` e `/personals`)

#### `GET /personals`
Retorna as informações de perfis pessoais armazenadas.
- **Resposta Sucesso (`200 OK`):**
  ```json
  [
    {
      "id_personal": 1,
      "address": "Rua Exemplo",
      "city": "São Paulo",
      "neighborhood": "Centro",
      "state": "SP",
      "cep": "01000-000",
      "phone": "(11) 99999-9999",
      "email": "contato@exemplo.com",
      "website": "https://meusite.com",
      "linkedin": "https://linkedin.com/in/usuario",
      "github": "https://github.com/usuario"
    }
  ]
  ```

---

## 🛠️ Contribuição
Sinta-se à vontade para enviar *pull requests* ou abrir *issues* com sugestões e melhorias.
