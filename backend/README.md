# Restaurant Orders Backend

Este é um sistema de gerenciamento de pedidos para restaurantes desenvolvido em Go, utilizando PostgreSQL para persistência e Server-Sent Events (SSE) para comunicação em tempo real entre garçons e a cozinha.

## 🚀 Tecnologias

- **Linguagem:** Go
- **Framework Web:** [Gin Gonic](https://gin-gonic.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Banco de Dados:** PostgreSQL
- **Ambiente:** Nix Flakes (Automação de dependências e banco local)
- **Comunicação:** SSE (Server-Sent Events)

## 🛠️ Configuração e Execução

O projeto utiliza **Nix Flakes** para configurar automaticamente o ambiente de desenvolvimento, incluindo o PostgreSQL local.

### 1. Entrar no ambiente
Na raiz do projeto, execute:
```bash
nix develop
```
*Isso configurará o Go e o PostgreSQL localmente na pasta `.db`.*

### 2. Gerenciar o Banco de Dados
Com o shell do Nix ativo, utilize os aliases configurados:

1.  **Iniciar o banco de dados:**
    ```bash
    pg-start
    ```
2.  **Acessar o terminal do Postgres:**
    ```bash
    pg-cli
    ```
3.  **Criar Tipos Customizados:**
    Como o GORM `AutoMigrate` não cria tipos `ENUM` automaticamente no PostgreSQL, você deve executar o script inicial uma vez para definir os tipos e as constraints:
    ```sql
    \i enums.sql
    ```
    *Após isso, o GORM gerenciará as tabelas automaticamente ao iniciar a aplicação.*

### 3. Rodar a aplicação
Certifique-se de que o banco de dados está ligado, instale as dependências e execute:
```bash
cd src/
go mod tidy
go run main.go
```
O servidor iniciará em `http://localhost:8080`.

## 📌 Endpoints Principais

- `POST /establesh/singin`: Autenticação e registro do estabelecimento.
- `POST /api/orders`: Criação de pedidos.
- `GET /api/stream/kitchen?establishment_id={id}`: Canal SSE para a cozinha receber pedidos em tempo real.

---
*Dica: Para desligar o banco ao terminar, use `pg-stop`.*