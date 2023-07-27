FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN chmod -x wait-for-db.sh
RUN chmod -x dev-startup.sh

EXPOSE 8080
CMD ["air"]