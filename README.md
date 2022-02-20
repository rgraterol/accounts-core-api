# Accounts Core API
API that handle money accounts core logic in an online wallet network.

![image](https://user-images.githubusercontent.com/2694731/154365849-a1a5b0a1-ba40-42e4-8c9b-715acd5981bb.png)



# Technologies
- Go 1.17
- MySQL
- Prometheus
- Grafana
- Docker - Alpine
- [WIP] Datadog


# To run the project
- Initialize all the containers, including the go docker image
```
docker-compose up -d
```

- Test if server is up and running
```bash
curl --location --request GET 'http://localhost:8080/ping'
```

**Notice: ZAP creates the folder `/logs` which creates `.txt` files in the `MM-DD-YYY`format to record logs**

# To DEBUG the project
- Comment the `go-app`, `prometheus`, and `grafana` containers in docker-compose.
- In the configuration file `development.yml` change the DB connection string to aim `localhost:3306`

## Tests

To test the application run the following command

````bash
go test  ./... -covermode=atomic  -coverpkg=./... -count=1  -race -timeout=30m
````

# Endpoints

## Users-API Feed consumer `POST /users`
Receives news from users-api and create new accounts on the registration of new users
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

## User account `GET /users/{userID}`
These endpoints retrieve the account for a given userID
```bash
curl --location --request GET 'localhost:8080/users/{userID}'
```

Returns an `Account` object of the type
```json
{
    "id": 1,
    "name": 1,
    "currency_id": "",
    "country": "",
    "available_amount": 1600,
    "block_reason": ""
}
```

## Deposit `POST /movements/deposit/3`
For a user to have money, he or she must have to deposit first. This can be done with the cURL
```bash
curl --location --request POST 'http://localhost:8080/movements/deposit/{userID}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 500,
    "currency_id": "EUR"
    "reason":"Account creation"
}'
```

#### Input
The input is a deposit object of the type.
- Amount: The specific amount to deposit.
- CurrencyID: The currency of the deposit.
- Reason: Small description of the deposit.
```json
{
    "amount": 500,
    "currency_id": "EUR",
    "reason":"Account creation"
}
```

#### Output
In a deposit the payer a collector are the same, but we still record it as a movement.
So we return a movement object like:
```json
{
    "id": 6,
    "payer_user_id": 3,
    "payer_account_id": 3,
    "collector_user_id": 3,
    "collector_account_id": 3,
    "amount": 500,
    "payer_balance": 1500,
    "collector_balance": 2000,
    "reason": "Account creation",
    "currency_id": "EUR",
    "country_id": "",
    "status": "done"
}
```


## Transfer `POST /movements/transfer`
To make transfers between accounts we use the transfer endpoint.
```bash
curl --location --request POST 'http://localhost:8080/movements/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "payer_user_id": 1,
    "collector_user_id": 2,
    "deposit": {
        "amount": 200,
        "currency_id": "EUR"
    }
}'
```

#### Input

- CollectorUserID: User to whom the money will be deposited
- PayerUserID: User who is making the transfer, to whom we will debit the money
- Deposit: The financial information of the transfer
- - Amount: The specific amount to deposit.
- - CurrencyID: The currency of the deposit.
- - Reason: Small description of the deposit.


#### Output
The output is a movement object in which the collector and payer are different.
```json
{
    "payer_user_id": 1,
    "payer_account_id": 1,
    "collector_user_id": 2,
    "collector_account_id": 2,
    "amount": 200,
    "payer_balance": 600,
    "collector_balance": 400,
    "reason": "",
    "currency_id": "EUR",
    "country_id": "",
    "status": "done"
}
```

# Fallbacks

If there's a problem writing to the DB while making a movement, the movement record on the DB is registered as failed
and the updates to the accounts made are rollback in a way that the users don't get failed debits or deposits.


# Play Around

1 - To play around the API we need first to create an account with the `POST /users` endpoint
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

2 - Then we need to deposit him some money with
```bash
curl --location --request POST 'http://localhost:8080/movements/deposit/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 500,
    "currency_id": "EUR"
    "reason":"Account creation"
}'
```
3 - We can check if we successfully made the deposit looking and his account 
```bash
curl --location --request GET 'localhost:8080/users/1'
```

4 - To transfer to another user we first need to create the other user with 
```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msg": {
        "id": 2,
        "headers": {
            "new_user": true
        }
    }
}'
```

5 - Now we can transfer the new user 200 EUR from the deposited 500 EUR
```bash
curl --location --request POST 'http://localhost:8080/movements/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "payer_user_id": 1,
    "collector_user_id": 2,
    "deposit": {
        "amount": 200,
        "currency_id": "EUR"
    }
}'
```

6 - Finally we check if the new user successfully received the transfer
```bash
curl --location --request GET 'localhost:8080/users/2'
```


# Killing feature
A nice feature that the API doesn't handle completely but it has some things already developed is a multi-site, multi-currency feature. For the users to transfer money between countries.
We now support the country and currency on the data objects but for this feature we might need to make an exchange rate service inside the API calling another API that gives us the exchange rate between currencies for us to calculate it.

# Folder Structure

```
...
├── configs 
│   └── Configuration `.yml` files
├── application // Layer that communicates the domain with the infra. Example: controllers, restclients, DBRepositories.
│   ├── db
│   ├── handlers
│   ├── middlewares
│   ├── repositories
│   ├── responses
│   └── router
├── domain // Business logic layer
│   ├── CheckoutController.java
│   └── PingController.java
└── infrastructure // Outer layer that connects with external services like DB, logging.
    ├── init
    └── main.go

...
```
