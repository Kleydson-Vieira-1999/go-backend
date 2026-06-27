# --- Estágio de Compilação (Build) ---
FROM docker.io/library/golang:1.26-alpine3.22 AS builder

# Instala dependências de build necessárias para CGO (se precisar)
RUN apk add --no-cache alpine-sdk

WORKDIR /app

# Copia os arquivos de dependências primeiro (aproveita o cache de camadas)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Compila o binário estático otimizado para produção
# CGO_ENABLED=0 garante que o binário não dependa de bibliotecas dinâmicas do host
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# --- Estágio de Execução (Runtime) ---
FROM docker.io/library/alpine:3.22

WORKDIR /app

# Cria um usuário não-root por segurança (essencial para boas práticas de SaaS)
RUN adduser -D restauranteuser

COPY --from=builder /app/client_secret_apps.googleusercontent.com.json ./

RUN chown restauranteuser:restauranteuser ./client_secret_apps.googleusercontent.com.json

# Copia apenas o binário compilado do estágio anterior
COPY --from=builder /app/main .

USER restauranteuser

# Porta padrão que o Gin vai escutar (ajuste se usar outra)
EXPOSE 8080 

ENV JWT_PASS=

CMD ["./main"]