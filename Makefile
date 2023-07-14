include .env
export $(shell sed 's/=.*//' .env)

#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up:
	migrate -path pkg/sql/migrations -database $$DB_URL  -verbose up

migrate-down:
	migrate -path pkg/sql/migrations -database $$DB_URL  -verbose down

migrate-force:
	migrate -path pkg/sql/migrations -database $$DB_URL  -verbose force $(version)

migrate-create:
	migrate create -ext sql -dir pkg/sql/migrations -seq $(name)
