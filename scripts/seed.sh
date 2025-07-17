#!/bin/bash

# Script para executar seeds do banco de dados
# Uso: ./scripts/seed.sh [seed|clear]

set -e

# Verificar se foi passado um argumento
if [ $# -eq 0 ]; then
    echo "Uso: $0 [seed|clear]"
    echo "  seed  - Popular o banco com dados fictícios"
    echo "  clear - Limpar todos os dados do banco"
    exit 1
fi

# Navegar para a pasta raiz do projeto
cd "$(dirname "$0")/.."

# Executar o comando
case $1 in
    "seed")
        echo "🌱 Executando seeds..."
        go run seeds/main.go seed
        ;;
    "clear")
        echo "🧹 Limpando banco de dados..."
        go run seeds/main.go clear
        ;;
    *)
        echo "Comando inválido. Use 'seed' ou 'clear'"
        exit 1
        ;;
esac 