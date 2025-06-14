# ✈️ Travel GO – Backend

Este é o backend da aplicação **Travel GO**, desenvolvido em **Golang** com **Fiber** e **PostgreSQL**. Ele fornece APIs para autenticação, gerenciamento de pedidos de viagem e notificações.

---

## 📦 Requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🚀 Como rodar o projeto (via Docker)

### 1. Clone o repositório

```bash
git clone https://github.com/igormpb/travel-go.git
cd travel-go
```

### 2. Rode os containers

```bash
docker-compose up --build
```

> O backend será iniciado em: [http://localhost:3002](http://localhost:3002)

---

## ⚙️ Configuração de ambiente

Você pode configurar variáveis de ambiente criando um arquivo `.env` na raiz do projeto:

```env
JWT_SECRET=supersecretjwt123
DATABASE_URL=postgres://postgres:postgres@db:5432/travel?sslmode=disable
```

> Em produção, essas variáveis devem ser configuradas diretamente no ambiente do container.

---

## 🧪 Como executar os testes

Caso você tenha testes implementados (em breve):

```bash
go test ./...
```

> Se estiver usando Docker:

```bash
docker exec -it travel-go-app-1 go test ./...
```


## 🐳 Docker

- Porta exposta: `3002`
- Banco de dados: `PostgreSQL 15`
- Dados persistidos em `volume: db_data`

---

## 📫 Contato

Desenvolvido por [Igor Barros](https://github.com/igormpb)  
