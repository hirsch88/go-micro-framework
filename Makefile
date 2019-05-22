include .env
export

install:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install

mock:
	@echo "Generating mocks..."
	mockery --version || go get -u -f github.com/vektra/mockery/.../
	mockery -dir=./app -recursive=true -all

test:
	@echo "Running unit tests..."
	go test -v ./app/...

run:
	@echo "Running application..."
	go run main.go

build:
	@echo "Building application..."
	go build -o out main.go

db-status:
	@echo "Migratin database..."
	goose -dir database/migrations mysql "$(DB_CONNECTION)" status

db-migrate:
	@echo "Migratin database..."
	goose -dir database/migrations mysql "$(DB_CONNECTION)" up

db-rollback:
	@echo "Migratin database..."
	goose -dir database/migrations mysql "$(DB_CONNECTION)" down

db-reset:
	@echo "Migratin database..."
	goose -dir database/migrations mysql "$(DB_CONNECTION)" reset
	goose -dir database/migrations mysql "$(DB_CONNECTION)" up

db-create-migration:
	@echo "Createing a new migration file..."
	goose -dir database/migrations create $(name) sql

