DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up:
 	migrate -path pkg/sql/migrations -database "postgresql://postgres:toor@172.168.0.2:5432/inventory_management?sslmode=disable"  -verbose up

migrate-down:
	migrate -path pkg/sql/migrations -database "postgresql://postgres:toor@localhost:5432/inventory_management?sslmode=disable"  -verbose down

migrate-force:
	migrate -path pkg/sql/migrations -database "postgresql://postgres:toor@localhost:5432/inventory_management?sslmode=disable"  -verbose force $(version)

migrate-create:
	migrate create -ext sql -dir pkg/sql/migrations -seq $(name)