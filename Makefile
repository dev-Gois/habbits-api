# Makefile para Habbits API

.PHONY: help run build seed clear-db clean install

# Comando padrão
help:
	@echo "Comandos disponíveis:"
	@echo "  make run         - Executar o servidor"
	@echo "  make build       - Compilar a aplicação"
	@echo "  make install     - Instalar dependências"
	@echo "  make seed        - Popular banco com dados fictícios"
	@echo "  make clear-db    - Limpar todos os dados do banco"
	@echo "  make clean       - Limpar arquivos compilados"

# Executar o servidor
run:
	go run main.go

# Compilar a aplicação
build:
	go build -o bin/habbits-api main.go

# Instalar dependências
install:
	go mod download
	go mod tidy

# Popular banco com dados fictícios
seed:
	./scripts/seed.sh seed

# Limpar dados do banco
clear-db:
	./scripts/seed.sh clear

# Limpar arquivos compilados
clean:
	rm -rf bin/ 