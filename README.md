# Widgets Single Page App API [![Build Status](https://travis-ci.org/danilojunS/widgets-spa-api.svg?branch=master)](https://travis-ci.org/danilojunS/widgets-spa-api)

> My try on the backend API for the widgets-spa project! ✌️

## Introduction

This is my implementation of the widgets-spa backend, written in Go.

This application basically consists on a CRUD for two concepts: **users** and **widgets**, exposed as an HTTP API authenticated via Token.

I tried to borrow some concepts from [DDD](https://en.wikipedia.org/wiki/Domain-driven_design) (Doman Drive Design). Some of the concepts used here are:
- entities: objects defined by its identity in the domain.
- use cases: what the application does (their business purpose). The use cases could be discussed in an [ubiquitous language](https://martinfowler.com/bliki/UbiquitousLanguage.html) with all tech and business stakeholders.

Thus, the project is organized in two main directories: 
- **business**: contains all business logic of the application. In this case, `entities`, `use cases` and `repositories`. A `repository` is the persistance layer for the `entities` and are an interface between the business and infra. Ideally, the code in the business layer could and should be discussed by all members (tech and business) in the organization, as they represent the domain business logic.
- **infra**: contains all logic of how the application is accessed. In this case, it contains the code to start a server and create a connection with the database. This layer could be easily replaced by another one, as it does not contain any business logic. For example, the CRUD could be accessed by CLI tool, instead of an HTTP API.

For illustration, the flow of a request is:

1. the `request` is routed to a `handler` (infra) of the server, which is responsible to call the appropriate `use case`.
1. the `use case` receives parameters of the `request` and:
    1. instantiates the appropriate `entities`.
    1. validates the `entitiy` (business logic involved).
    1. calls the necessary `services`.
    1. processes the data and apply any business logic.
    1. persists the data, using a `repository`.
    1. returns result to `handler`.
1. the `handler` receives the result from the `use case`, builds the response and send it to the client.

With the appropriate separation between the business and infra, we can build applications whose business logic can be re-used. Moreover, it promotes applications more aligned to their purpose in the organization, as both business and tech teams can discuss it and improve it more seamlessly. 🙌

The next sections describe how to install, run and test the application.

## Installing and running

### With Docker 🐳 (recommended)

These steps assume that you have [docker](https://docs.docker.com/engine/installation/) installed in your computer.

1. Get the repository
  
    ```shell
    go get -v github.com/danilojunS/widgets-spa-api
    ```
    
1. Go to the repository directory and run
  
    ```shell
    make docker-compose-start
    ```
    
    This command uses [docker-compose](https://docs.docker.com/compose/) to build and start two containers: one for the application (`4000` port) and another for the database (`5432` port).
    
1. [optional] Populate the DB container with dummy data
  
    ```shell
    make docker-compose-db-populate
    ```

And that's it! 👍

You can test the API in the url: `http://localhost:4000` (view [Try it](#try-it) section).

### Without Docker

This steps assume that you have [Go](https://golang.org/doc/install) and [Postgres](https://wiki.postgresql.org/wiki/Detailed_installation_guides) installed and configured in your computer.

1. Get the repository

    ```shell
    go get -v github.com/danilojunS/widgets-spa-api
    ```
    
1. Install the dependencies

    ```shell
    make install
    ```
    
1. Start the db

    ```shell
    pg_ctl -D /path/to/db/data -l logfile start
    ```

1. [optional] Populate the DB (the command uses the default `postgres` database)

    ```shell
    make db-populate
    ```
1. Start the application

    ```shell
    make start
    ```
    
## Other usefull commands

All commands are declared and documented in the `Makefile`. 

Some other useful commands (and their docker version) are:

- `make test` or `make docker-test`
  
  Run the automated tests.
  
  The docker version of this command requires to run the `make docker-build` command first.
  
- `make watch` or `make docker-compose-watch`
  
  Run the application in `watch` mode, that is, every change in the source codes will automatically restart your application.

## Endpoints

This application exposes the following endpoints:

- GET `/token` [http://localhost:4000/token](http://localhost:4000/token)
- 🔒 GET `/users` [http://localhost:4000/users](http://localhost:4000/users)
- 🔒 GET `/users/:id` [http://localhost:4000/users/:id](http://localhost:4000/users/:id)
- 🔒 GET `/widgets` [http://localhost:4000/widgets](http://localhost:4000/widgets)
- 🔒 GET `/widgets/:id` [http://localhost:4000/widgets/:id](http://localhost:4000/widgets/:id)
- 🔒 POST `/widgets` for creating new widgets [http://localhost:4000/widgets](http://localhost:4000/widgets)
- 🔒 PUT `/widgets/:id` for updating existing widgets [http://localhost:4000/widgets/:id](http://localhost:4000/widgets/:id)

The 🔒 endpoints needs authentication using a token obtained in the `GET /token` endpoint.

The tokens must be passed in the `Authorization` header.

Eg.: `Authorization: Bearer <TOKEN>`

## <a name="try-it"></a>Try it

After running it in you local machine, you can use this Postman Collection to easily try the API.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/24e9d2eeb998c28834be#?env%5Bwidgets-spa-dev%5D=W3siZW5hYmxlZCI6dHJ1ZSwia2V5IjoidXJsIiwidmFsdWUiOiJsb2NhbGhvc3Q6NDAwMCIsInR5cGUiOiJ0ZXh0In1d)

Or use the good old `curl`:

- Get token
  ```
  curl -X GET \
    http://localhost:4000/token
  ```
- Get users
  ```
  curl -X GET \
    http://localhost:4000/users \
    -H 'authorization: Bearer <TOKEN>'
  ```
- Get users/{id}
  ```
  curl -X GET \
    http://localhost:4000/users/1 \
    -H 'authorization: Bearer <TOKEN>'
  ```
- Get widgets
  ```
  curl -X GET \
    http://localhost:4000/widgets \
    -H 'authorization: Bearer <TOKEN>'
  ```
- Get widgets/{id}
  ```
  curl -X GET \
    http://localhost:4000/widgets/1 \
    -H 'authorization: Bearer <TOKEN>'
  ```
- Create widgets
  ```
  curl -X POST \
    http://localhost:4000/widgets \
    -H 'authorization: Bearer <TOKEN>' \
    -H 'content-type: application/json' \
    -d '{
      "name": "My new widget",
      "color": "red",
      "price": "19.99",
      "inventory": 42,
      "melts": true
    }'
  ```
- Update widgets
  ```
  curl -X PUT \
    http://localhost:4000/widgets/1 \
    -H 'authorization: Bearer <TOKEN>' \
    -H 'content-type: application/json' \
    -d '{
      "name": "My other widget",
      "color": "red",
      "price": "19.99",
      "inventory": 42,
      "melts": true
    }'
  ```

## Improvements

Although this application was fun to build and test, there is still some improvements that could be done (implement end-to-end tests, increase test coverage, refactors to avoid repetition, etc).

Also, this was my first Go application 🙈
Some parts of the code are not as idiomatic as a seasoned Go programmer would write.

Please, feel free to open up an issue/PR if you feel like discussing these aspects!
