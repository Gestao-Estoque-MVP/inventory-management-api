include .env
export $(shell sed 's/=.*//' .env)

#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up:
	migrate -path pkg/sql/migrations -database $$DB_URL_MIGRATE  -verbose up


migrate-up-version:
	migrate -path pkg/sql/migrations -database $$DB_URL_MIGRATE -verbose up $(version)

migrate-down:
	migrate -path pkg/sql/migrations -database $$DB_URL_MIGRATE  -verbose down

migrate-force:
	migrate -path pkg/sql/migrations -database $$DB_URL_MIGRATE  -verbose force $(version)

migrate-create:
	migrate create -ext sql -dir pkg/sql/migrations -seq $(name)
