FROM golang:1.17-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache gcc
RUN apk add --no-cache sqlite
RUN apk add libc-dev

ADD . /go/src/accounts-core-api

RUN mv /go/src/accounts-core-api /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install /app/infrastructure/init/main.go

RUN chmod +x /app/infrastructure/init/main.go

EXPOSE 8080

ENTRYPOINT ["/app/infrastructure/init/main.go"]

RUN echo "Server running on port 8080"