# Include
include .env

# Migrations
migration-create:
	migrate create -ext sql -dir db/migrations -seq $(name) 

migration-up:
	@echo "$(CYAN)Starting migration... $(GREEN) UP $(NC)"
	migrate -path db/migrations -database ${DB_URL} -verbose up

migration-down:
	@echo "$(CYAN)Starting migration... $(RED) DOWN $(NC)"
	migrate -path db/migrations -database ${DB_URL} -verbose down 

# Seeds
seeds-run:
	@echo "$(GREEN)Seeding database... 🌱 $(NC)"
	go run seeds/main.go 
