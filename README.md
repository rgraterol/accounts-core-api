# Transfers Core API
API that handle money transfers core logic between accounts in an online wallet network.

# Technologies
- Go 1.17
- MySQL


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