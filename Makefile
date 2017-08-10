# Production

install:
	go get -v
.PHONY: install

build:
	go build main.go

clean:
	rm main
.PHONY: clean

start-prod: build
	APP_ENV=production
	./main
.PHONY: start-prod

# Development

start-dev:
	APP_ENV=development
	go run main.go
.PHONY: start-dev

test:
	APP_ENV=testing
	go test -cover ./...
.PHONY: test

# Data

db-populate:
	psql -d postgres -a -f scripts/sql/populate-users-widgets.sql
