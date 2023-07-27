#!/bin/sh

# Run Goose with the provided command and arguments
echo "RUN MIGRATE:"
goose -dir db/migrations postgres "postgres://postgres:P@ssword@postgres:5432/prl?sslmode=disable" up


# Run application
echo "API START"
air 