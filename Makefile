include .env

export DATABASE_URL ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

bin:
	@mkdir -p bin

setup-tools: bin
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
ifeq ($(shell uname), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
	echo "Downloaded"
else ifeq ($(shell uname), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

formatting:
	go fmt ./...

migration-create:
	bin/migrate create -dir migrations -tz utc -ext sql $(name)

migration-up:
	bin/migrate -path migrations -database "${DATABASE_URL}" up

migration-down:
	bin/migrate -path migrations -database "${DATABASE_URL}" down $(n)

seed-create:
	bin/migrate create -ext sql -dir migrations/seeds -seq $(name)

seed-up:
	bin/migrate -path migrations/seeds -database "${DATABASE_URL}&x-migrations-table=seed_migrations" up

seed-down:
	bin/migrate -path migrations/seeds -database "${DATABASE_URL}&x-migrations-table=seed_migrations" down $(n)

run-dev:
	bin/air -d --build.cmd "go build -o tmp/api main.go" --build.bin "tmp/api"

run-rest:
	bin/air -d --build.cmd "go build -o tmp/rest cmd/rest/rest.go" --build.bin "tmp/rest"

.PHONY: 
	formatting 
	gen-doc
	run-test
	migration-create
	migration-up
	migration-down
	seed-create
	seed-up
	seed-down
	run-dev
	run-rest