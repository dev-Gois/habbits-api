# Habbits API

Uma API RESTful desenvolvida em Go para gerenciamento de hábitos e rotinas diárias.

## 📋 Descrição

A Habbits API é uma aplicação backend construída com Go, Gin e GORM que permite aos usuários gerenciar seus hábitos, definir metas e acompanhar seu progresso diário. A API oferece funcionalidades de autenticação JWT, criação de usuários, gerenciamento de hábitos e controle de check-ins diários.

## 🚀 Tecnologias Utilizadas

- **Go 1.24.4** - Linguagem de programação
- **Gin** - Framework web para Go
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação e autorização
- **bcrypt** - Criptografia de senhas
- **godotenv** - Gerenciamento de variáveis de ambiente
- **validator** - Validação de dados

## 📁 Estrutura do Projeto

```
habbits-api/
├── config/
│   └── database.go          # Configuração do banco de dados
├── controllers/
│   ├── application_controller.go  # Controlador de aplicação
│   └── users_controller.go       # Controlador de usuários
├── models/
│   ├── user.go              # Modelo de usuário
│   ├── habit.go             # Modelo de hábito
│   └── habit_check.go       # Modelo de check-in de hábito
├── routes/
│   └── router.go            # Definição das rotas
├── services/
│   ├── jwt/
│   │   ├── encode.go        # Geração de tokens JWT
│   │   └── decode.go        # Decodificação de tokens JWT
│   └── users/
│       └── create.go        # Lógica de criação de usuários
├── go.mod                   # Dependências do Go
├── go.sum                   # Checksums das dependências
└── main.go                  # Ponto de entrada da aplicação
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
   JWT_SECRET=sua_chave_secreta_jwt
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

| Método | Endpoint | Descrição | Autenticação |
|--------|----------|-----------|--------------|
| GET | `/` | Endpoint de teste da aplicação | Não |
| POST | `/register` | Criar novo usuário | Não |

### Exemplos de Uso

#### Criar Usuário
```bash
curl -X POST http://localhost:3000/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "senha123"
  }'
```

**Resposta:**
```json
{
  "message": "Usuário criado com sucesso!",
  "user": {
    "id": 1,
    "name": "João Silva",
    "email": "joao@example.com",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 🗄️ Modelos de Dados

### User
- `id` - ID único do usuário
- `name` - Nome do usuário (mínimo 3 caracteres)
- `email` - Email do usuário (único)
- `password` - Senha criptografada
- `habits` - Relacionamento com hábitos do usuário

### Habit
- `id` - ID único do hábito
- `title` - Título do hábito
- `icon` - Ícone do hábito
- `sunday` - Habilitado para domingo
- `monday` - Habilitado para segunda
- `tuesday` - Habilitado para terça
- `wednesday` - Habilitado para quarta
- `thursday` - Habilitado para quinta
- `friday` - Habilitado para sexta
- `saturday` - Habilitado para sábado
- `user_id` - ID do usuário proprietário
- `checks` - Relacionamento com check-ins

### HabitCheck
- `id` - ID único do check-in
- `habit_id` - ID do hábito
- `done` - Status de conclusão
- `date` - Data do check-in

## 🔐 Autenticação

A API utiliza JWT (JSON Web Tokens) para autenticação:

- **Geração de Token**: Ao criar um usuário, um token JWT é gerado automaticamente
- **Validade**: Tokens são válidos por 30 dias
- **Segurança**: Senhas são criptografadas usando bcrypt

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
- `github.com/golang-jwt/jwt/v5` - Autenticação JWT
- `golang.org/x/crypto/bcrypt` - Criptografia de senhas
- `github.com/go-playground/validator/v10` - Validação de dados

## 🚧 Funcionalidades Implementadas

- ✅ Configuração do banco de dados PostgreSQL
- ✅ Modelos de dados (User, Habit, HabitCheck)
- ✅ Autenticação JWT
- ✅ Criação de usuários com validação
- ✅ Criptografia de senhas
- ✅ Estrutura de rotas básica

## 🔄 Próximas Funcionalidades

- [ ] Login de usuários
- [ ] CRUD completo de hábitos
- [ ] Sistema de check-ins diários
- [ ] Middleware de autenticação
- [ ] Relatórios e estatísticas
- [ ] Notificações

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