# TaxCalc
William Minar Widjaja

## Getting started

### Clone repository

`$ git clone github.com/filifolia/TaxCalc.git`


## Running in development

### Prerequisites

You need the following:

- Docker
- Docker-compose
- Go installed

# Start the application

Run $ docker-compose up -d

# Database migration

In-case the database is not automatically migrated, please run this command:
$ docker-compose exec api bee migrate -driver=postgres -conn="postgres://postgres:password@postgres/tax-db?sslmode=disable"

# Setting up test environment

Create a test database.

$ docker-compose exec postgres createdb -h postgres -p 5432 -U postgres -E UTF8 tax-db-test

Input the db password "password" when prompted

# Test Database migration

$ docker-compose exec api bee migrate -driver=postgres -conn="postgres://postgres:password@postgres/tax-db-test?sslmode=disable"

# Unit test 
Remove the "-e /models" tag if you want to include database models integration test as well

If you are in UNIX environment:
$ docker-compose exec api env TAX_ENV=test go test -cover $(go list ./... | grep -v -e /vendor -e /database/migrations -e /models -e tests)

If you are in Windows, you can run this commands instead in order:
$ docker-compose exec api bash
$ env TAX_ENV=test go test -cover $(go list ./... | grep -v -e /vendor -e /database/migrations -e /models -e tests)

# Access the API

Once the server has been started, you can start using the API from localhost:8080

Example: $ curl localhost:8080/v1/item/get