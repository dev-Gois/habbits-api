# Seeds do Banco de Dados

Este diretório contém scripts para popular o banco de dados com dados fictícios para demonstração e desenvolvimento.

## Como usar

### Opção 1: Script automatizado (recomendado)
```bash
# Popular o banco com dados fictícios
./scripts/seed.sh seed

# Limpar todos os dados do banco
./scripts/seed.sh clear
```

### Opção 2: Executar manualmente
```bash
# Navegar para a pasta seeds
cd seeds

# Popular o banco
go run main.go seed

# Limpar o banco
go run main.go clear
```

## O que os seeds criam

### Usuários (5 usuários)
- **Pedro Feijó** - feijo@gmail.com
- **Sextou da Cantina** - sextou@gmail.com
- **Germano Fenner** - germano@gmail.com
- **Douglas Saboia** - saboia@gmail.com
- **Jose Henrique** - henrique@gmail.com
- Senha padrão: `123456` (hash bcrypt)

### Hábitos (8 hábitos por usuário)
- **Variedade**: Água, exercício, meditação, leitura, programação, caminhada, sono, frutas, alongamento, diário, gratidão, organização, vitaminas, redes sociais, yoga
- **Ícones**: Emojis apropriados para cada hábito
- **Dias da semana**: Diferentes padrões (todos os dias, dias úteis, fins de semana, dias alternados, etc.)

### Habit Checks (últimos 60 dias)
- **Padrões realistas**: Cada hábito tem uma taxa de completude específica
  - Tomar vitaminas: 90% (mais fácil)
  - Beber água: 85% (rotina básica)
  - Praticar gratidão: 75% (moderado)
  - Estudar programação: 70% (dedicação)
  - Dormir 8 horas: 65% (desafio comum)
  - Exercitar-se: 60% (requer disciplina)
  - Ler: 55% (requer tempo)
  - Fazer alongamento: 50% (esquecimento comum)
  - Meditar: 45% (difícil manter)
  - Caminhar 10k passos: 40% (meta alta)
  - Fazer yoga: 40% (requer prática)
  - Escrever diário: 35% (requer constância)
  - Organizar ambiente: 30% (tarefa pesada)
  - Evitar redes sociais: 25% (muito difícil)

- **Horários realistas**: Distribuição baseada em padrões de comportamento
  - Manhã (6h-9h): Maior concentração
  - Tarde/Noite (16h-23h): Segundo pico
  - Horários aleatórios dentro dos períodos

## Benefícios para demonstração

1. **Gráficos populados**: Dashboard com dados suficientes para visualizações
2. **Diferentes estados**: Hábitos completos, parciais e não feitos
3. **Padrões realistas**: Dados que simulam comportamento real do usuário
4. **Variedade temporal**: 60 dias de histórico para análises
5. **Múltiplos usuários**: Demonstrar isolamento de dados por usuário

## Tecnologias utilizadas

- **GORM**: ORM para criação em massa
- **Bcrypt**: Hash seguro de senhas
- **Time**: Manipulação de datas e horários
- **Math/rand**: Geração de padrões aleatórios realistas

## Estrutura dos dados

```
Users (5)
├── Habits (8 cada = 40 total)
│   ├── Diferentes padrões de dias da semana
│   ├── Ícones e nomes variados
│   └── Taxas de completude específicas
└── HabitChecks (baseado em 60 dias)
    ├── Apenas para dias válidos do hábito
    ├── Horários realistas
    └── Padrões de completude variados
```

## Limpeza

O comando `clear` remove todos os dados das tabelas:
- `habit_checks` (primeiro, devido a foreign keys)
- `habits` 
- `users`

⚠️ **Atenção**: O comando clear remove TODOS os dados do banco, incluindo dados reais se existirem. 