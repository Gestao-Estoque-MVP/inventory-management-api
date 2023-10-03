# Etapa de compilação
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o api cmd/main.go

# Etapa de criação da imagem final
FROM alpine
COPY --from=builder /app/api /api
EXPOSE 80
ENTRYPOINT [ "/api" ]
