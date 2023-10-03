# Etapa de compilação
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o api cmd/main.go

# Etapa de criação da imagem final
FROM alpine
COPY --from=builder /app/api /api
EXPOSE 8080
ENTRYPOINT [ "/api" ]
