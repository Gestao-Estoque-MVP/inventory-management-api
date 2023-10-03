# Etapa de compilação
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o api cmd/main.go

# Etapa de criação da imagem final
FROM debian:bullseye-slim

# Instale quaisquer dependências que você possa precisar. 
# Por exemplo, `ca-certificates` é útil para chamadas HTTPS.
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/api /api
RUN chmod +x /api

EXPOSE 8080
ENTRYPOINT [ "/api" ]