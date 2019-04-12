#!/bin/bash

bee migrate -driver=postgres -conn="postgres://postgres:password@postgres/tax-db?sslmode=disable"