# Usa uma imagem base de Go 1.21 (ou 1.22, se necessário)
FROM golang:1.22-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Instala pacotes adicionais para compilar a aplicação
RUN apk add --no-cache gcc musl-dev

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências da aplicação
RUN go mod download

# Executa go mod tidy para ajustar dependências no go.mod
RUN go mod tidy

# Copia o código da aplicação para o diretório de trabalho
COPY . .

# Compila a aplicação
RUN go build -o go-app

# Cria uma imagem mais leve a partir do Alpine
FROM alpine:3.18

# Define o diretório de trabalho para a imagem final
WORKDIR /root/

# Copia o binário da etapa de build para o container final
COPY --from=builder /app/go-app .

# Define a porta que o container vai expor
EXPOSE 8000

# Comando para rodar a aplicação
CMD ["./go-app"]