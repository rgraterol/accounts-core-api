FROM golang:1.17.7 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app/ ./...

EXPOSE 8080

CMD ["/app/accounts-core-api"]

RUN echo "Server running on port 8080"