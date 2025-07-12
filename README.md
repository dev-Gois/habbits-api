# Habbits API

Uma API RESTful desenvolvida em Go para gerenciamento de hábitos e rotinas diárias.

## 📋 Descrição

A Habbits API é uma aplicação backend construída com Go, Gin e GORM que permite aos usuários gerenciar seus hábitos, definir metas e acompanhar seu progresso diário. A API oferece funcionalidades de autenticação JWT, criação e login de usuários, gerenciamento de hábitos e controle de check-ins diários.

## 🚀 Tecnologias Utilizadas

- **Go 1.24.4** - Linguagem de programação
- **Gin** - Framework web para Go
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação e autorização
- **bcrypt** - Criptografia de senhas
- **godotenv** - Gerenciamento de variáveis de ambiente
- **validator** - Validação de dados
- **CORS** - Cross-Origin Resource Sharing configurado

## 📁 Estrutura do Projeto

```
habbits-api/
├── config/
│   └── database.go          # Configuração do banco de dados
├── controllers/
│   ├── application_controller.go  # Controlador de aplicação
│   ├── users_controller.go       # Controlador de usuários
│   └── habits_controller.go      # Controlador de hábitos
├── middlewares/
│   └── authorization.go     # Middleware de autorização
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
│   ├── users/
│   │   ├── create.go        # Lógica de criação de usuários
│   │   └── login.go         # Lógica de login de usuários
│   ├── habits/
│   │   ├── create.go        # Lógica de criação de hábitos
│   │   ├── find_all.go      # Lógica de busca de hábitos
│   │   ├── find.go          # Lógica de busca de hábito específico
│   │   ├── update.go        # Lógica de atualização de hábitos
│   │   ├── delete.go        # Lógica de exclusão de hábitos
│   │   ├── verify_weekday_on_create.go  # Verificação de dia da semana na criação
│   │   └── get_habit_day_value.go       # Obter valor do dia do hábito
│   ├── habit_checks/
│   │   ├── create.go        # Lógica de criação de check-ins 
│   │   ├── handle_today.go  # Lógica de manipulação do dia atual
│   │   └── delete_incomplete.go  # Lógica de exclusão de check-ins incompletos
│   └── cron/
│       └── scheduler.go     # Scheduler de cron jobs automatizados
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
| GET | `/api` | Endpoint de teste da aplicação | Não |
| POST | `/api/register` | Criar novo usuário | Não |
| POST | `/api/login` | Fazer login de usuário | Não |
| GET | `/api/user` | Obter dados do usuário logado | Sim |
| POST | `/api/habits` | Criar novo hábito | Sim |
| GET | `/api/habits` | Obter todos os hábitos do usuário | Sim |
| PUT | `/api/habits/:id` | Atualizar um hábito específico | Sim |
| DELETE | `/api/habits/:id` | Deletar um hábito específico | Sim |
| GET | `/api/habit-checks` | Obter todos os check-ins do dia do usuário | Sim |
| PUT | `/api/habit-checks/:id/check` | Marcar/desmarcar check-in como concluído | Sim |
| POST | `/api/workers/create-habit-checks` | Executar job de criar check-ins para uma data | Não |

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

#### Fazer Login
```bash
curl -X POST http://localhost:3000/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@example.com",
    "password": "senha123"
  }'
```

**Resposta:**
```json
{
  "message": "Login realizado com sucesso!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Obter Dados do Usuário (Autenticado)
```bash
curl -X GET http://localhost:3000/api/user \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
{
  "message": "Usuário encontrado com sucesso!",
  "user": {
    "id": 1,
    "name": "João Silva",
    "email": "joao@example.com",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

#### Criar Hábito (Autenticado)
```bash
curl -X POST http://localhost:3000/api/habits \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Exercitar-se",
    "icon": "🏃‍♂️",
    "sunday": true,
    "monday": true,
    "tuesday": true,
    "wednesday": true,
    "thursday": true,
    "friday": true,
    "saturday": false
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "title": "Exercitar-se",
  "icon": "🏃‍♂️",
  "sunday": true,
  "monday": true,
  "tuesday": true,
  "wednesday": true,
  "thursday": true,
  "friday": true,
  "saturday": false,
  "user_id": 1,
  "user": {
    "id": 1,
    "name": "João Silva",
    "email": "joao@example.com",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### Obter Todos os Hábitos (Autenticado)
```bash
curl -X GET http://localhost:3000/api/habits \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
[
  {
    "id": 1,
    "title": "Exercitar-se",
    "icon": "🏃‍♂️",
    "sunday": true,
    "monday": true,
    "tuesday": true,
    "wednesday": true,
    "thursday": true,
    "friday": true,
    "saturday": false,
    "user_id": 1,
    "user": {
      "id": 1,
      "name": "João Silva",
      "email": "joao@example.com",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    "checks": [],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  {
    "id": 2,
    "title": "Ler 30 minutos",
    "icon": "📚",
    "sunday": false,
    "monday": true,
    "tuesday": true,
    "wednesday": true,
    "thursday": true,
    "friday": true,
    "saturday": true,
    "user_id": 1,
    "user": {
      "id": 1,
      "name": "João Silva",
      "email": "joao@example.com",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    "checks": [],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
]
```

#### Deletar Hábito (Autenticado)
```bash
curl -X DELETE http://localhost:3000/api/habits/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
{
  "message": "Hábito deletado com sucesso"
}
```

**Erro (Hábito não encontrado):**
```json
{
  "error": "hábito não encontrado"
}
```

#### Atualizar Hábito (Autenticado)
```bash
curl -X PUT http://localhost:3000/api/habits/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Comer Brownie",
    "icon": "🍫",
    "sunday": true,
    "monday": true,
    "tuesday": true,
    "wednesday": true,
    "thursday": true,
    "friday": true,
    "saturday": true
  }'
```

**Resposta:**
```json
{
  "id": 1,
  "title": "Comer Brownie",
  "icon": "🍫",
  "sunday": true,
  "monday": true,
  "tuesday": true,
  "wednesday": true,
  "thursday": true,
  "friday": true,
  "saturday": true,
  "user_id": 5,
  "user": {
    "id": 5,
    "name": "João Silva",
    "email": "joao@example.com",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "checks": [],
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

#### Obter Check-ins do Dia (Autenticado)
```bash
curl -X GET http://localhost:3000/api/habit-checks \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
[
  {
    "id": 1,
    "habit_id": 1,
    "done": false,
    "date": "2024-06-07T00:00:00Z",
    "habit": {
      "id": 1,
      "title": "Exercitar-se",
      "icon": "🏃‍♂️",
      "user_id": 1
    }
  },
  {
    "id": 2,
    "habit_id": 2,
    "done": true,
    "date": "2024-06-07T00:00:00Z",
    "habit": {
      "id": 2,
      "title": "Ler 30 minutos",
      "icon": "📚",
      "user_id": 1
    }
  }
]
```

#### Marcar/Desmarcar Check-in (Autenticado)
```bash
curl -X PUT http://localhost:3000/api/habit-checks/1/check \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
{
  "message": "Check-in realizado com sucesso"
}
```

**Erro (Check-in não encontrado):**
```json
{
  "error": "check-in não encontrado"
}
```

#### Executar Job Manualmente (Criar Check-ins)
```bash
# Para hoje
curl -X POST http://localhost:3000/api/workers/create-habit-checks

# Para uma data específica (query parameter)
curl -X POST "http://localhost:3000/api/workers/create-habit-checks?date=2024-06-07"

# Para uma data específica (JSON body)
curl -X POST http://localhost:3000/api/workers/create-habit-checks \
  -H "Content-Type: application/json" \
  -d '{"date": "2024-06-07"}'
```

**Resposta:**
```json
{
  "message": "Job executado com sucesso",
  "date": "2024-06-07",
  "created": 3
}
```

**Erro (Data inválida):**
```json
{
  "error": "Formato de data inválido. Use YYYY-MM-DD"
}
```

## 🗄️ Modelos de Dados

### User
- `id` - ID único do usuário
- `name` - Nome do usuário (mínimo 3 caracteres)
- `email` - Email do usuário (único)
- `password` - Senha criptografada
- `plain_password` - Senha em texto plano (apenas para input)
- `habits` - Relacionamento com hábitos do usuário

**Métodos:**
- `SetPassword(password)` - Criptografa e define a senha
- `CheckPassword(password)` - Verifica se a senha está correta
- `Create()` - Cria o usuário com validações
- `Get()` - Busca o usuário pelo ID

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
- `user` - Dados do usuário proprietário
- `checks` - Relacionamento com check-ins

**Métodos:**
- `Create()` - Cria o hábito

### HabitCheck
- `id` - ID único do check-in
- `habit_id` - ID do hábito
- `done` - Status de conclusão
- `date` - Data do check-in

## 🔐 Autenticação

A API utiliza JWT (JSON Web Tokens) para autenticação:

- **Geração de Token**: Ao criar um usuário ou fazer login, um token JWT é gerado automaticamente
- **Validade**: Tokens são válidos por 30 dias
- **Segurança**: Senhas são criptografadas usando bcrypt
- **Decodificação**: Função para decodificar tokens e extrair o ID do usuário
- **Middleware**: Middleware de autorização para proteger rotas

### Serviços JWT

#### Encode (Geração de Token)
```go
token, err := jwt.Encode(userID)
```

#### Decode (Decodificação de Token)
```go
userID, err := jwt.Decode(tokenString)
```

### Middleware de Autorização

O middleware `Authorization()` verifica:
- Presença do header `Authorization`
- Formato correto: `Bearer <token>`
- Validade do token JWT
- Existência do usuário no banco de dados

## 🔧 Desenvolvimento

### Executar em modo de desenvolvimento
```bash
go run main.go
```

### Build da aplicação
```bash
go build -o habbits-api main.go
```

### Configuração CORS

A API está configurada para aceitar requisições de qualquer origem localhost:

- **AllowOrigins**: `*` (aceita qualquer origem)
- **AllowMethods**: GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
- **AllowHeaders**: Origin, Content-Length, Content-Type, Authorization
- **AllowCredentials**: true
- **MaxAge**: 12 horas

Isso permite que frontends rodando em diferentes portas (3000, 3001, 8080, etc.) acessem a API sem problemas de CORS.

## 📦 Dependências Principais

- `github.com/gin-gonic/gin` - Framework web
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - Driver PostgreSQL
- `github.com/joho/godotenv` - Gerenciamento de variáveis de ambiente
- `github.com/golang-jwt/jwt/v5` - Autenticação JWT
- `golang.org/x/crypto/bcrypt` - Criptografia de senhas
- `github.com/go-playground/validator/v10` - Validação de dados
- `github.com/gin-contrib/cors` - Middleware CORS

## 🚧 Funcionalidades Implementadas

- ✅ Configuração do banco de dados PostgreSQL
- ✅ Modelos de dados (User, Habit, HabitCheck)
- ✅ Autenticação JWT (encode/decode) - **CORRIGIDO**
- ✅ Criação de usuários com validação
- ✅ Login de usuários
- ✅ Criptografia de senhas com bcrypt
- ✅ Validação de dados com validator
- ✅ Middleware de autorização - **CORRIGIDO**
- ✅ Estrutura de rotas básica
- ✅ Tratamento de erros personalizado
- ✅ Endpoint protegido `/user`
- ✅ Criação de hábitos - **NOVO**
- ✅ Busca de hábitos do usuário - **NOVO**
- ✅ Exclusão de hábitos - **NOVO**
- ✅ Atualização de hábitos com regras especiais - **NOVO**
- ✅ Endpoint para buscar check-ins do dia (`GET /habit-checks`) - **NOVO**
- ✅ Endpoint para marcar/desmarcar check-ins (`PUT /habit-checks/:id/check`) - **NOVO**
- ✅ Sistema de cron jobs automatizados - **NOVO**
- ✅ Endpoint para executar jobs manualmente - **NOVO**

## ⏰ Cron Jobs Automatizados

A API inclui um sistema de cron jobs que executa tarefas automatizadas:

### Jobs Configurados

| Cron Expression | Descrição | Frequência |
|----------------|-----------|------------|
| `0 0 * * *` | Criar check-ins diários | Todo dia à 00:00 |

### Funcionalidades

- **Criação Automática de Check-ins**: Cria automaticamente check-ins para hábitos ativos do dia

### Como Funciona

1. **Check-ins Diários**: Todo dia à meia-noite, o sistema verifica todos os hábitos ativos e cria check-ins para os que devem ser executados naquele dia da semana

### Execução Manual

Você também pode executar jobs manualmente através do endpoint:
- **POST** `/api/workers/create-habit-checks` - Cria check-ins para uma data específica
  - **Query Parameter**: `date` (opcional, formato: YYYY-MM-DD)
  - **JSON Body**: `{"date": "YYYY-MM-DD"}` (opcional)
  - **Sem data**: Executa para hoje
  - **Com data**: Executa para a data especificada

## 🐛 Correções Recentes

### Sistema JWT
- **Problema**: Token inválido devido a incompatibilidade de tipos
- **Solução**: Convertido `userID` para `float64` na codificação e `float64` para `uint` na decodificação
- **Melhoria**: Adicionado melhor tratamento de erros e validação de expiração

### Middleware de Autorização
- **Problema**: Mensagens de erro genéricas
- **Solução**: Mensagens de erro mais específicas e informativas
- **Melhoria**: Adicionado contexto adicional (`userID`) para uso posterior

### Modelo User
- **Problema**: Método `Get()` não funcionava corretamente
- **Solução**: Corrigido para buscar pelo ID correto do usuário

### Criação de Hábitos
- **Problema**: Falta de endpoint para criação de hábitos
- **Solução**: Implementado endpoint `POST /habits` com autenticação
- **Melhoria**: Hábito criado automaticamente associado ao usuário logado

### Busca de Hábitos
- **Problema**: Falta de endpoint para buscar hábitos do usuário
- **Solução**: Implementado endpoint `GET /habits` com autenticação
- **Melhoria**: Retorna todos os hábitos do usuário logado com dados completos

### Exclusão de Hábitos
- **Problema**: Endpoint de delete não verificava propriedade do hábito e não retornava resposta adequada
- **Solução**: Implementado verificação de propriedade e tratamento de erros específicos
- **Melhoria**: Agora verifica se o hábito existe e pertence ao usuário antes de deletar

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