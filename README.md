# Deck API

Simple deck management REST API

- [Deck API](#deck-api)
  - [Language](#language)
  - [Used libraries](#used-libraries)
  - [How to run](#how-to-run)
  - [How to test](#how-to-test)
  - [Docs](#docs)

## Language

This project is written in _Golang_ 1.22

## Used libraries

- [Cobra-cli](https://github.com/spf13/cobra)
- [Echo](https://github.com/labstack/echo)
- [Logrus](https://github.com/sirupsen/logrus)

## How to run

The project can run locally but first will need to

1. Download and install [Go](https://go.dev/)
2. Download this repository

   ```bash
    git clone https://github.com/jvfrodrigues/transaction-product-wex
   ```

3. Install all dependencies
   ```bash
   go mod download
   ```
4. Run the project
   ```bash
   go run main.go
   #OR
   make serve
   ```
5. Test! If no changes were made it will run [localhost:8484](http://localhost:8484) with Swagger docs available at [localhost:8484/swagger/index.html](http://localhost:8484/swagger/index.html)

## How to test

1. You can run tests by
   ```bash
   go test ./...
   #OR
   make test
   #OR for coverage
   make coverage
   ```

## Docs

Docs are available through Swagger Open API Documentation at [localhost:8484/swagger/index.html](http://localhost:8484/swagger/index.html) when server is running
