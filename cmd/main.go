package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/factory"
	"github.com/diogoX451/inventory-management-api/internal/graph"
	"github.com/diogoX451/inventory-management-api/internal/graph/directives"
	"github.com/diogoX451/inventory-management-api/internal/graph/middleware"
	"github.com/diogoX451/inventory-management-api/internal/repository"
	"github.com/diogoX451/inventory-management-api/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func init() {

	if err := godotenv.Load(); err != nil {
		panic("No .env variable")
	}

}



func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v\n", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v\n", err)
		}
	}()

	logFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de log: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	queries := database.New(db)
	rcba := repository.NewRBCARepository(queries)
	rcbaService := service.NewRCBAService(rcba)
	s3Repository := repository.NewS3Repository(queries)
	templateRepository := repository.NewTemplateRepository(queries)
	userRepository := repository.NewRepositoryUser(queries)
	fact := factory.NewSendEmailFactory(templateRepository, userRepository, configS3(s3Repository))
	userService := service.NewServiceUser(userRepository, rcba, fact)
	contactRepository := repository.NewRepositoryContactInfo(queries)
	contactService := service.NewContactInfoService(contactRepository)
	loginService := service.NewAuthUser(*userRepository, *rcba)
	addressRepository := repository.NewAddressRepository(queries)
	addressRepositoryService := service.NewAddressService(addressRepository)

	service.NewTemplateService(*templateRepository)

	resolvers := &graph.Resolver{
		UserRepository:        userRepository,
		UserService:           userService,
		ContactInfoRepository: contactRepository,
		ContactInfoService:    contactService,
		RBCARepository:        rcba,
		RBCAService:           rcbaService,
		AuthUserService:       loginService,
		AddressRepository:     addressRepository,
		AddressService:        addressRepositoryService,
		S3Repository:          *s3Repository,
		S3Service:             configS3(s3Repository),
	}

	c := graph.Config{
		Resolvers: resolvers,
		Directives: graph.DirectiveRoot{
			Auth:    directives.Auth,
			HasRole: directives.HasRole,
		},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
