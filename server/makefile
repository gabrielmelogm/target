# Include
include .env

# Migrations
migration-create:
	migrate create -ext sql -dir db/migrations -seq $(name) 

migration-up:
	@echo "$(CYAN)Starting migration... $(GREEN) UP $(NC)"
	migrate -path db/migrations -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}${DB_SSL_MODE}" -verbose up

migration-down:
	@echo "$(CYAN)Starting migration... $(RED) DOWN $(NC)"
	migrate -path db/migrations -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}${DB_SSL_MODE}" -verbose down 

# Seeds
seeds-run:
	@echo "$(GREEN)Seeding database... 🌱 $(NC)"
	go run seeds/main.go 
