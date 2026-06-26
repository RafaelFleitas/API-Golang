# API-Golang

API REST de gerenciamento de usuários desenvolvida em Go, com Oracle Database como banco de dados.

## Tecnologias

| Pacote | Versão | Descrição |
|--------|--------|-----------|
| Go | 1.21+ | Linguagem |
| [Gin](https://github.com/gin-gonic/gin) | v1.12.0 | Framework HTTP |
| [go-ora](https://github.com/sijms/go-ora) | v2.9.0 | Driver Oracle puro Go (sem Oracle Instant Client) |
| [godotenv](https://github.com/joho/godotenv) | v1.5.1 | Carregamento de variáveis de ambiente |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | — | Hash de senhas |
| [Zap](https://go.uber.org/zap) | v1.28.0 | Logger estruturado |

## Arquitetura

```
main.go
src/
├── configuration/
│   ├── database/oracleSQL/   # Conexão com Oracle
│   ├── logger/               # Logger com Zap
│   ├── rest_err/             # Padronização de erros HTTP
│   └── validation/           # Validação de campos de entrada
├── controller/
│   ├── model/
│   │   ├── request/          # Structs de entrada (JSON → Go)
│   │   └── response/         # Structs de saída (Go → JSON)
│   └── routes/               # Registro das rotas
├── model/
│   ├── repository/           # Acesso ao banco de dados
│   └── service/              # Casos de uso / regras de negócio
└── view/                     # Conversão domínio → response
```

### Fluxo de uma requisição

```
HTTP Request → Controller → Service → Repository → Oracle DB
                                                        ↓
HTTP Response ← View (converte) ← Service ←────────────┘
```

## Endpoints

| Método | Rota | Descrição | Status |
|--------|------|-----------|--------|
| `POST` | `/createUser` | Cria novo usuário | Implementado |
| `GET` | `/getUserById/:userId` | Busca usuário por ID | Implementado |
| `GET` | `/getUserByEmail/:userEmail` | Busca usuário por e-mail | Implementado |
| `PUT` | `/updateUser/:userId` | Atualiza nome e/ou idade | Implementado |
| `DELETE` | `/deleteUser/:userId` | Remove usuário | Em desenvolvimento |

---

### POST `/createUser`

**Request body:**

```json
{
  "name": "Rafael Fleitas",
  "email": "rafael@email.com",
  "password": "Senha@123",
  "age": 25
}
```

**Validações:**

| Campo | Regra |
|-------|-------|
| `name` | Obrigatório, entre 4 e 50 caracteres |
| `email` | Obrigatório, formato de e-mail válido |
| `password` | Obrigatório, mínimo 6 caracteres, deve conter ao menos um de: `! @ # $ % *` |
| `age` | Obrigatório, entre 1 e 140 |

**Response `200 OK`:**

```json
{
  "id": 1,
  "name": "Rafael Fleitas",
  "email": "rafael@email.com",
  "age": 25
}
```

> A senha **nunca** é retornada. Ela é armazenada como hash bcrypt.

---

### GET `/getUserById/:userId`

**Parâmetro de rota:** `userId` — inteiro numérico.

**Response `200 OK`:**

```json
{
  "id": 1,
  "name": "Rafael Fleitas",
  "email": "rafael@email.com",
  "age": 25
}
```

**Erros possíveis:**

| Código | Motivo |
|--------|--------|
| `400` | `userId` não é um número inteiro válido |
| `404` | Usuário não encontrado |

---

### GET `/getUserByEmail/:userEmail`

**Parâmetro de rota:** `userEmail` — endereço de e-mail válido.

**Response `200 OK`:**

```json
{
  "id": 1,
  "name": "Rafael Fleitas",
  "email": "rafael@email.com",
  "age": 25
}
```

**Erros possíveis:**

| Código | Motivo |
|--------|--------|
| `400` | E-mail com formato inválido |
| `404` | Usuário não encontrado |

---

### PUT `/updateUser/:userId`

**Parâmetro de rota:** `userId` — inteiro numérico.

Permite atualizar `name` e/ou `age`. Ambos os campos são opcionais; ao menos um deve ser enviado.

**Request body:**

```json
{
  "name": "Novo Nome",
  "age": 30
}
```

**Validações:**

| Campo | Regra |
|-------|-------|
| `name` | Opcional, entre 4 e 50 caracteres, não pode ser vazio |
| `age` | Opcional, entre 1 e 140 |

**Response `200 OK`:**

```json
{
  "id": 1,
  "name": "Novo Nome",
  "email": "rafael@email.com",
  "age": 30
}
```

**Erros possíveis:**

| Código | Motivo |
|--------|--------|
| `400` | `userId` inválido ou campos fora das regras de validação |
| `404` | Usuário não encontrado |

---

### Formato de erro

Todos os erros seguem a mesma estrutura:

```json
{
  "message": "Some fields are invalid",
  "error": "bad_request",
  "code": 400,
  "causes": [
    {
      "field": "password",
      "message": "password must contain at least one special character"
    }
  ]
}
```

## Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) e Docker Compose

## Como rodar

### 1. Subir o banco de dados Oracle

```bash
docker compose up -d
```

Aguarde o container ficar com status `healthy` (pode levar 2–3 minutos na primeira vez):

```bash
docker ps
```

O Oracle será exposto na porta `1521`.

### 2. Configurar variáveis de ambiente

Copie o arquivo de exemplo e preencha com suas configurações:

```bash
cp .env.example .env
```

Edite o `.env`:

```env
ORACLE_URL=oracle://user:password@localhost:port/service
ORACLE_PASSWORD=PASSWORD_HERE
```

Substitua `user`, `password` e `service` pelas suas credenciais Oracle.

### 3. Instalar dependências

```bash
go mod tidy
```

### 4. Rodar a aplicação

```bash
go run main.go
```

A API sobe na porta padrão do Gin (`8080`).

## Banco de dados

A tabela de usuários esperada pelo repositório:

```sql
CREATE TABLE users (
  id       NUMBER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name     VARCHAR2(50)  NOT NULL,
  email    VARCHAR2(255) NOT NULL,
  password VARCHAR2(255) NOT NULL,
  age      NUMBER(3)     NOT NULL
);
```
