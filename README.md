# API-Golang

## Tecnologias

- **Go 1.26**
- **Gin** — framework HTTP
- **Oracle Database (Free)** — banco de dados via Docker
- **go-ora** — driver Oracle puro Go (sem Oracle Instant Client)
- **godotenv** — carregamento de variáveis de ambiente
- **bcrypt** — hash de senhas
- **Zap** — logger estruturado

## Arquitetura

```
main.go
src/
├── configuration/
│   ├── database/oracleSQL/   # Conexão com Oracle
│   ├── logger/               # Logger com Zap
│   ├── rest_err/             # Padronização de erros HTTP
│   └── validation/           # Validação de campos
├── controller/
│   ├── model/
│   │   ├── request/          # Structs de entrada (JSON → Go)
│   │   └── response/         # Structs de saída (Go → JSON)
│   └── routes/               # Registro das rotas
├── model/                    # Domínio (regras de negócio)
│   └── service/              # Casos de uso
└── view/                     # Conversão domínio → response
```

### Fluxo de uma requisição

```
HTTP Request → Controller → Service → Model (domínio)
                                            ↓
HTTP Response ← View (converte) ←──────────┘
```

## Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/getUserById/:userId` | Busca usuário por ID |
| GET | `/getUserByEmail/:userEmail` | Busca usuário por e-mail |
| POST | `/createUser` | Cria novo usuário |
| PUT | `/updateUser/:userId` | Atualiza usuário |
| DELETE | `/deleteUser/:userId` | Remove usuário |

## Pré-requisitos

- [Go](https://golang.org/dl/) 1.21+
- [Docker](https://www.docker.com/) e Docker Compose

## Como rodar

### 1. Subir o banco de dados Oracle

```bash
docker compose up -d
```

Aguarde o container ficar com status `healthy` (pode levar 2-3 minutos na primeira vez):

```bash
docker ps
```

### 2. Configurar variáveis de ambiente

Renomeie o arquivo `.env.example` para `.env` e preencha com as informações.

### 3. Instalar dependências

```bash
go mod tidy
```

### 4. Rodar a aplicação

```bash
go run main.go
```