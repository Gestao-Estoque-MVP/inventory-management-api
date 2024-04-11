package main

import (
	"context"
	"log"
	"os"

	companies_router "github.com/diogoX451/inventory-management-api/cmd/router/companies"
	users_router "github.com/diogoX451/inventory-management-api/cmd/router/users"
	"github.com/diogoX451/inventory-management-api/internal/database"
	middlewares "github.com/diogoX451/inventory-management-api/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

}

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
	}

	defer func() {
		db.Close()
	}()

	logFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de log: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	router := gin.Default()
	api := router.Group("/")

	autorize := router.Group("/api/v1")
	autorize.Use(middlewares.Auth())

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	queries := database.New(db)

	users_router.RouterUsers(queries, autorize, api)
	companies_router.RouterCompanies(queries, autorize)

	router.Run()

	// s3Repository := repository.NewS3Repository(queries)
	// image := repository.NewImageRepository(queries)
	// s3Service := service.NewServiceS3(s3Repository, os.Getenv("S3_BUCKET_NAME"), os.Getenv("S3_ACESS_KEY_ID"), os.Getenv("S3_REGION"))
	// imageService := service.NewImageService(*image, s3Service)
	// rcba := repository.NewRBCARepository(queries)
	// rcbaService := service.NewRCBAService(rcba)
	// userRepository := repository.NewRepositoryUser(queries)
	// emailService := service.NewServiceEmail(s3Repository, *userRepository)
	// userService := service.NewServiceUser(userRepository, rcba, emailService, s3Service)
	// contactRepository := repository.NewRepositoryContactInfo(queries)
	// contactService := service.NewContactInfoService(contactRepository, emailService)
	// loginService := service.NewAuthUser(*userRepository, *rcba)
	// addressRepository := repository.NewAddressRepository(queries)
	// addressRepositoryService := service.NewAddressService(addressRepository)

}
