Documentação da API Go-API
Descrição
Este é um projeto de uma aplicação web simples implementada em Go, utilizando o framework Gin para criar uma API RESTful. A API permite gerenciar produtos através de operações básicas como listagem, criação e busca por ID.

Estrutura do Projeto
Dependências
Gin: Framework web para Go.
Gorm: ORM (Object Relational Mapping) para Go.
Arquivos Principais
main.go: Ponto de entrada da aplicação, onde o servidor Gin é configurado e as rotas são definidas.
controller/product_controller.go: Define os controladores (controllers) para manipulação dos endpoints relacionados aos produtos.
repository/product_repository.go: Implementa o repositório para interações com o banco de dados.
usecase/product_usecase.go: Define os casos de uso (use cases) para a lógica de negócios dos produtos.
db/database.go: Configura e estabelece conexão com o banco de dados.
Como Executar
Clonar o repositório:

git clone https://github.com/seu-usuario/go-api.git
cd go-api
Instalar Dependências:

go mod download
Configurar o Banco de Dados:

Renomear config_example.yml para config.yml.
Configurar as variáveis de ambiente no arquivo config.yml conforme necessário.
Executar a Aplicação:

go run main.go
Endpoints Disponíveis
Health Check
GET /ping
Descrição: Retorna um "pong" para verificar se o servidor está online.
Resposta:
{
  "message": "pong"
}
Produtos
GET /products

Descrição: Lista todos os produtos disponíveis.
Resposta: Array de objetos JSON representando os produtos.
POST /product

Descrição: Cria um novo produto.
Corpo da Requisição (JSON):
{
  "name": "Nome do Produto",
  "description": "Descrição do Produto"
}
Resposta: Objeto JSON representando o produto criado.
GET /product/:productId

Descrição: Retorna os detalhes de um produto específico pelo ID.
Parâmetro na URL: productId (ID do produto).
Resposta: Objeto JSON representando o produto.
Estrutura das Rotas e Controladores
Controller
ProductController é responsável por definir as rotas relacionadas aos produtos.
func NewProductController(uc usecase.ProductUseCase) *ProductController {
    return &ProductController{useCase: uc}
}
Repository
ProductRepository abstrai as operações de banco de dados relacionadas aos produtos.
func NewProductRepository(db *gorm.DB) repository.ProductRepository {
    return &productRepository{db: db}
}
Use Case
ProductUseCase contém a lógica de negócios para operações com produtos.
func NewProductUseCase(repo repository.ProductRepository) usecase.ProductUseCase {
    return &productUseCase{repository: repo}
}
Considerações Finais
Este projeto é um exemplo básico de uma API RESTful em Go, utilizando Gin para rotas e Gorm para interação com o banco de dados. Para fins educacionais, demonstra como estruturar um projeto Go com múltiplos componentes interconectados.

Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests para melhorias e correções.

Espero que isso ajude a documentar seu projeto Go API de forma clara e útil!