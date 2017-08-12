# MAKEFILE: WIDGETS SPA API

####################################################################
## ENVIRONMENT
##
## Definition of default values for some environment variables.
## You can overwrite them by simply setting the env variable
## before running the commands in this Makefile.
## Eg.: `export APP_ENV=production`
####################################################################

# Application environment: production, development, testing
APP_ENV ?= "development"
# Port of the webserver
PORT ?= 4000

####################################################################
## MAIN
##
## Commands to natively run the application.
## They assume that you have go and postgres installed in your
## computer.
## Alternatively, if you wish to run the application with docker,
## you can run the commands in the `DOCKER` and `DOCKER-COMPOSE`
## sessions above.
## Commands: install, build, clean, test, start and watch
####################################################################

# Install the application dependencies and packages
.PHONY: install
install:
	go get -v
	go get -v github.com/codegangsta/gin

# Build the application executable
build:
	go build main.go

# Clean the files generated in the build process
.PHONY: clean
clean:
	go clean

# Run the automated tests
.PHONY: test
test:
	export APP_ENV=testing; go test -cover ./...

# Start the application (does not require build)
.PHONY: start
start:
	go run main.go

# Watch files and restart the app when any of them change
.PHONY: watch
watch:
	gin --port ${PORT} run main.go

####################################################################
## DOCKER
##
## Commands to run the application inside docker.
## These commands simply start the application inside a container,
## assuming that you already have setup the db in your computer.
## If you do not have a running instance of postgres,
## please use the docker-compose commands above.
####################################################################

# Build the docker image with the application
.PHONY: docker-build
docker-build:
	docker build -t widgets-spa-api .

# Start a docker container with the application image
.PHONY: docker-start
docker-start:
	docker run \
		-p ${PORT}:${PORT} \
		-e APP_ENV=${APP_ENV} \
		widgets-spa-api make start

# Start a docker container with the application image in watch mode
# Any changes in the application files will restart the application
.PHONY: docker-watch
docker-watch:
	docker run \
		-p ${PORT}:${PORT} \
		-v `pwd`:/go/src/github.com/danilojunS/widgets-spa-api \
		-e APP_ENV=${APP_ENV} \
		widgets-spa-api make watch

# Run the automated tests in the docker container
.PHONY: docker-test
docker-test:
	docker run \
		widgets-spa-api make test

####################################################################
## DOCKER-COMPOSE
##
## Commands to setup a docker environment with the application
## and the database.
## It generates one container for the application and another
## for the database, linking them together.
## Use them if you wish to have the db inside a container,
## when you do not have postgres installed in your computer.
## Obs.: there is no test script, as the tests does not require
## a running database.
####################################################################

# Start:
# 1. Docker container with the application image.
# 2. Docker container with the postgres database.
.PHONY: docker-compose-start
docker-compose-start:
	export COMMAND="make start"; docker-compose up

# Start:
# 1. Docker container with the application image in watch mode.
# 	 Any changes in the application files will restart the application.
# 2. Docker container with the postgres database.
.PHONY: docker-compose-watch
docker-compose-watch:
	export COMMAND="make watch"; docker-compose up

# Kills the containers created by the `make docker-compose-start` and
# `make docker-compose-watch` commands
.PHONY: docker-compose-kill
docker-compose-kill:
	docker-compose kill

# Populate the database inside the container with dummy data
.PHONY: docker-compose-db-populate
docker-compose-db-populate:
	docker-compose exec postgres psql -d postgres -U postgres -a -f /var/scripts/populate-users-widgets.sql

####################################################################
## POSTGRES
##
## Commands to perform actions in the database.
## They assume that you have a running instance of postgres in
## your computer
####################################################################

# Populate the database with dummy data
.PHONY: db-populate
db-populate:
	psql -d postgres -a -f scripts/sql/populate-users-widgets.sql
