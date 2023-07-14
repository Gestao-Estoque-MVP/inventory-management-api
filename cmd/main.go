package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/diogoX451/inventory-management-api/internal/database"
	"github.com/diogoX451/inventory-management-api/internal/graph"
	"github.com/diogoX451/inventory-management-api/internal/repositories"
	"github.com/diogoX451/inventory-management-api/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

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
	// Define a duração máxima que uma conexão pode ser reutilizada.
	db.SetConnMaxLifetime(time.Minute * 3)
	// Define um número máximo de conexões simultâneas.
	db.SetMaxOpenConns(10)
	// Define um número máximo de conexões inativas.
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

	port := os.Getenv("PORT")

	queries := database.New(db)

	userRepository := repositories.NewRepositoryUser(queries)
	userService := service.NewServiceUser(userRepository)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserRepository: userRepository,
		UserService:    userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
