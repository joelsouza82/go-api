# Documentação da API Go-API

## 📖 Descrição
Este projeto é uma API RESTful desenvolvida em **Go**, projetada utilizando **Arquitetura em Camadas (Clean Architecture)**. A aplicação oferece um CRUD completo para o gerenciamento de **informações pessoais/perfil** e de **credenciais de login**.

---

## 🛠️ Linguagem e Tecnologias Utilizadas

- **Linguagem:** [Go (Golang)](https://go.dev/) (v1.22+)
- **Framework Web:** [Gin Gonic](https://github.com/gin-gonic/gin) — Framework HTTP de alta performance para Go.
- **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/) (Hospedado na plataforma Supabase).
- **Driver de Banco de Dados:** `github.com/lib/pq` (Consultas SQL nativas através da biblioteca standard `database/sql`).
- **Testes:** `testify` (asserts e mocks) para testes unitários de controllers, use cases e repositories.
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
1. **Controller (`controller/`)**: Intercala a comunicação HTTP com os Casos de Uso. Formata as respostas JSON e trata códigos de status HTTP (200, 201, 204, 400, 404, 500).
2. **UseCase (`usecase/`)**: Implementa a lógica de negócios e as regras da aplicação.
3. **Repository (`repository/`)**: Abstrai o acesso aos dados e interage diretamente com o banco PostgreSQL por meio de queries SQL (`SELECT`, `INSERT`, `UPDATE`, `DELETE`).
4. **Model (`model/`)**: Define as estruturas de dados (structs) mapeadas em JSON.
5. **Mocks (`mocks/`)**: Contém os mocks das interfaces para utilização nos testes unitários.
6. **DB (`db/`)**: Gerencia e estabelece a conexão com a base de dados PostgreSQL.

---

## 🗄️ Banco de Dados

A API conecta-se a um banco de dados **PostgreSQL** utilizando consultas SQL nativas (`database/sql` + driver `github.com/lib/pq`).

### Estrutura das Tabelas:

#### Tabela `personal`
- `id` (INTEGER / SERIAL - Chave primária)
- `name`, `rg`, `document` (VARCHAR)
- `address`, `city`, `neighborhood`, `state`, `cep`, `phone`, `email`, `website`, `linkedin`, `github` (VARCHAR)
- `birthdate` (TIMESTAMP)
- `login_id` (INTEGER - Chave estrangeira para `login.id`)

#### Tabela `login`
- `id` (INTEGER / SERIAL - Chave primária)
- `email`, `password` (VARCHAR)

---

## 📂 Estrutura do Projeto

```text
go-api/
├── cmd/
│   └── main.go                        # Ponto de entrada e registro de rotas
├── controller/
│   ├── login_controller.go            # Controller para a entidade Login
│   ├── login_controller_test.go       # Testes do controller de Login
│   ├── personal_controller.go         # Controller para a entidade Personal
│   └── personal_controller_test.go    # Testes do controller de Personal
├── db/
│   └── conn.go                        # Conexão com o banco de dados
├── mocks/
│   ├── login_repository_mock.go       # Mock do repositório Login
│   ├── login_usecase_mock.go          # Mock do use case Login
│   ├── personal_repository_mock.go    # Mock do repositório Personal
│   └── personal_usecase_mock.go       # Mock do use case Personal
├── model/
│   ├── login.go                       # Struct do modelo Login
│   ├── personal.go                    # Struct do modelo Personal
│   └── response.go                    # Struct genérica de resposta
├── repository/
│   ├── login_repository.go            # Implementação do repositório de Login
│   ├── login_repository_interface.go  # Interface do repositório de Login
│   ├── login_repository_test.go       # Testes do repositório de Login
│   ├── personal_repository.go         # Implementação do repositório de Personal
│   ├── personal_repository_interface.go # Interface do repositório de Personal
│   └── personal_repository_test.go    # Testes do repositório de Personal
├── usecase/
│   ├── login_usecase.go               # Regras de negócio da entidade Login
│   ├── login_usecase_interface.go     # Interface do use case de Login
│   ├── login_usecase_test.go          # Testes do use case de Login
│   ├── personal_usecase.go            # Regras de negócio da entidade Personal
│   ├── personal_usecase_interface.go  # Interface do use case de Personal
│   └── personal_usecase_test.go       # Testes do use case de Personal
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

### 🧪 Executando os Testes

O projeto conta com testes unitários para os controllers, use cases e repositories, utilizando mocks das interfaces (pacote `mocks/`).

```bash
go test ./...
```

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

### 👤 Informações Pessoais (`/personal` e `/personals`)

#### `GET /personals`
Retorna todas as informações de perfis pessoais armazenadas.
- **Resposta Sucesso (`200 OK`):**
  ```json
  [
    {
      "id_personal": 1,
      "name": "Fulano de Tal",
      "rg": "12.345.678-9",
      "document": "123.456.789-00",
      "address": "Rua Exemplo",
      "city": "São Paulo",
      "neighborhood": "Centro",
      "state": "SP",
      "cep": "01000-000",
      "phone": "(11) 99999-9999",
      "email": "contato@exemplo.com",
      "website": "https://meusite.com",
      "linkedin": "https://linkedin.com/in/usuario",
      "github": "https://github.com/usuario",
      "birthdate": "1990-01-01T00:00:00Z",
      "login_id": 1
    }
  ]
  ```

#### `GET /personal/:personalId`
Retorna um perfil pessoal específico pelo ID.
- **Resposta Sucesso (`200 OK`):** objeto `Personal` (ver formato acima).
- **Resposta Erro (`400 Bad Request`):** ID inválido.
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

#### `POST /personal`
Cria um novo perfil pessoal.
- **Corpo da Requisição:** objeto `Personal` (sem `id_personal`).
- **Resposta Sucesso (`201 Created`):** objeto `Personal` criado.

#### `PUT /personal/:personalId`
Atualiza um perfil pessoal existente.
- **Corpo da Requisição:** objeto `Personal` com os campos a atualizar.
- **Resposta Sucesso (`200 OK`):** objeto `Personal` atualizado.
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

#### `DELETE /personal/:personalId`
Remove um perfil pessoal pelo ID.
- **Resposta Sucesso (`204 No Content`)**
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

---

### 🔐 Login (`/login` e `/logins`)

#### `GET /logins`
Retorna todos os registros de login armazenados.
- **Resposta Sucesso (`200 OK`):**
  ```json
  [
    {
      "id": 1,
      "email": "contato@exemplo.com",
      "password": "senha-hash"
    }
  ]
  ```

#### `GET /login/:loginId`
Retorna um registro de login específico pelo ID.
- **Resposta Sucesso (`200 OK`):** objeto `Login` (ver formato acima).
- **Resposta Erro (`400 Bad Request`):** ID inválido.
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

#### `POST /login`
Cria um novo registro de login.
- **Corpo da Requisição:** objeto `Login` (sem `id`).
- **Resposta Sucesso (`201 Created`):** objeto `Login` criado.

#### `PUT /login/:loginId`
Atualiza um registro de login existente.
- **Corpo da Requisição:** objeto `Login` com os campos a atualizar.
- **Resposta Sucesso (`200 OK`):** objeto `Login` atualizado.
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

#### `DELETE /login/:loginId`
Remove um registro de login pelo ID.
- **Resposta Sucesso (`204 No Content`)**
- **Resposta Erro (`404 Not Found`):** registro não encontrado.

---

## 🛠️ Contribuição
Sinta-se à vontade para enviar *pull requests* ou abrir *issues* com sugestões e melhorias.
