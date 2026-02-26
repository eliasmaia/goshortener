# 🚀 Go URL Shortener

Um encurtador de URLs ultra-rápido e minimalista construído em **Go**, utilizando **SQLite** para persistência e **Chi** para roteamento.

Este projeto demonstra os pilares fundamentais do ecossistema Go: performance, binários estáticos e tratamento explícito de erros.

## 🛠️ Tecnologias Utilizadas

* **Language:** Go 1.25+
* **Router:** [go-chi/chi](https://github.com/go-chi/chi) (Leve e compatível com net/http)
* **Database:** SQLite3 (Persistência em arquivo, sem necessidade de Docker para o banco)
* **Ambiente:** WSL2 (Ubuntu)

## 📋 Funcionalidades

- [x] Geração de códigos aleatórios de 6 caracteres.
- [x] Persistência de dados em SQLite.
- [x] Redirecionamento automático via HTTP 301.
- [x] API baseada em JSON.

## 🚀 Como Executar

### Pré-requisitos
* Go instalado (versão 1.21 ou superior).
* GCC (para o driver do SQLite). No Ubuntu/WSL: `sudo apt install build-essential`.

### Instalação
1. Clone o repositório:
   ```bash
   git clone https://github.com/eliasmaia/goshortener.git
   cd go-shortener

2. Instale as dependências:
    ```bash
    go mod tidy

3. Execute o projeto:
    ```bash
    go run .
## 🧪 Como Testar
1. Criar um link encurtado (POST)
    ```bash
    curl -X POST -H "Content-Type: application/json" \
    -d '{"url": "[https://www.google.com](https://www.google.com)"}' \
    http://localhost:8080/shorten
    ```
2. Acessar o redirecionamento (navegador)
Abra no seu navegador: 
    ``` http://localhost:8080/{codigo_gerado} ```

## 📁 Estrutura do Projeto
- main.go: Entrada da aplicação e handlers da API.
- db.go: Configuração e inicialização do SQLite.
- go.mod: Definição do módulo e dependências.

Desenvolvido por Elias Maia