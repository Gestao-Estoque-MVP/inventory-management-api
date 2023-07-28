# Etapa de compilação
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o api cmd/main.go

# Etapa de criação da imagem final
FROM alpine
COPY --from=builder /app/api /api

ENV DB_URL="postgresql://trombetasalomao:jkg0UGvtD2ss52quZWXf@swiftstock-db.cfbb4u4eeocz.sa-east-1.rds.amazonaws.com:5432/swiftstock"


ENTRYPOINT [ "/api" ]
