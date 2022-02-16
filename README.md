# Accounts Core API
API that handle money accounts core logic in an online wallet network.

![image](https://user-images.githubusercontent.com/2694731/154365540-e79375d3-b5a2-42f0-83e7-6c9f28137726.png)


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
