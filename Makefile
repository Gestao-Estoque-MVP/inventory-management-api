DOCKER_COMPOSE_FILE ?= docker-compose.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migrate-up:
	docker compose -f docker-compose.yml --profile tools run --rm migrate up

migrate-down:
	docker compose -f docker-compose.yml --profile tools run --rm migrate down

migrate-force:
	docker compose -f docker-compose.yml --profile tools run --rm migrate force $(version)

migrate-create:
	docker compose -f docker-compose.yml --profile tools run --rm migrate create -ext sql -dir /migrations $(name)