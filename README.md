# Habbits API

Uma API RESTful desenvolvida em Go para gerenciamento de hábitos e rotinas diárias.

## 📋 Descrição

A Habbits API é uma aplicação backend construída com Go, Gin e GORM que permite aos usuários gerenciar seus hábitos, definir metas e acompanhar seu progresso diário.

## 🚀 Tecnologias Utilizadas

- **Go 1.24.4** - Linguagem de programação
- **Gin** - Framework web para Go
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **godotenv** - Gerenciamento de variáveis de ambiente

## 📁 Estrutura do Projeto

```
habbits-api/
├── config/
│   └── database.go      # Configuração do banco de dados
├── controllers/
│   └── application_controller.go  # Controladores da aplicação
├── models/              # Modelos de dados
├── routes/
│   └── router.go       # Definição das rotas
├── services/           # Lógica de negócio
├── go.mod              # Dependências do Go
├── go.sum              # Checksums das dependências
└── main.go             # Ponto de entrada da aplicação
```

## 🛠️ Pré-requisitos

- Go 1.24.4 ou superior
- PostgreSQL
- Git

## ⚙️ Instalação

1. **Clone o repositório**
   ```bash
   git clone https://github.com/dev-Gois/habbits-api.git
   cd habbits-api
   ```

2. **Instale as dependências**
   ```bash
   go mod download
   ```

3. **Configure as variáveis de ambiente**
   
   Crie um arquivo `.env` na raiz do projeto:
   ```env
   DB_HOST=localhost
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_NAME=habbits_db
   DB_PORT=5432
   ```

4. **Configure o banco de dados**
   
   Certifique-se de que o PostgreSQL está rodando e crie o banco de dados:
   ```sql
   CREATE DATABASE habbits_db;
   ```

5. **Execute a aplicação**
   ```bash
   go run main.go
   ```

A API estará disponível em `http://localhost:3000`

## 📡 Endpoints

### Base URL
```
http://localhost:3000/api
```

### Endpoints Disponíveis

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/` | Endpoint de teste da aplicação |

## 🔧 Desenvolvimento

### Executar em modo de desenvolvimento
```bash
go run main.go
```

### Executar testes
```bash
go test ./...
```

### Build da aplicação
```bash
go build -o habbits-api main.go
```

## 📦 Dependências Principais

- `github.com/gin-gonic/gin` - Framework web
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - Driver PostgreSQL
- `github.com/joho/godotenv` - Gerenciamento de variáveis de ambiente

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👨‍💻 Autor

**Matheus Gois**
- GitHub: [@dev-Gois](https://github.com/dev-Gois)

## 📞 Suporte

Se você encontrar algum problema ou tiver dúvidas, abra uma issue no repositório.

---

⭐ Se este projeto te ajudou, considere dar uma estrela no repositório! 