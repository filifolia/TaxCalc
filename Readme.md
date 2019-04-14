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

# Unit test 

If you are in UNIX environment:
$ docker-compose exec api go test -cover $(go list ./... | grep -v -e /vendor -e /database/migrations -e /models -e tests)

If you are in Windows, you can run this commands instead in order:
$ docker-compose exec api bash
$ go test -cover $(go list ./... | grep -v -e /vendor -e /database/migrations -e /models -e tests)
