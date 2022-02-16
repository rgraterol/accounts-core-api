# Accounts Core API
API that handle money accounts core logic in an online wallet network.

![image](https://user-images.githubusercontent.com/2694731/154365849-a1a5b0a1-ba40-42e4-8c9b-715acd5981bb.png)



# Technologies
- Go 1.17
- MySQL
- Prometheus
- Grafana
- Docker - Alpine


# To run the project
- Initialize the DB
```
docker-compose up -d
```

- Download dependencies
```bash
go mod download
```

- Run Project
```bash
go run infrastructure/init/main.go
```

- Test if server is up and running
```bash
curl --location --request GET 'http://localhost:8080/ping'
``` 

## Tests

To test the application run the following command

````bash
go test  ./... -covermode=atomic  -coverpkg=./... -count=1  -race -timeout=30m
````

# Endpoints

## Users-API Feed consumer
Recieves news from users-api and create new accounts on the registration of new users
```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msg": {
        "id": 1,
        "headers": {
            "new_user": true
        }
    }
}'
```

## Movements
Even tho the /movements endpoint is in development, we can find all of it's business logic inside the domain layer `/domain/entities/movements.go` `/domain/interfacesmovements.go` and `/domain/services/movements.go`

cURL WIP

