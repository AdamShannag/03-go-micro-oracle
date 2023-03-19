# 03-go-micro-oracle

a microservice app made in golang that connects to an oracle database, using docker compose to manage the services.

## How to run
from the project directory, run the folowing
`sudo docker compose up` or `sudo docker-compose up`

then wait for the services to start.

## Endpoints
* > GET `http://localhost:8080/api/users`
* > POST `http://localhost:8080/api/users/id`
* > PUT`http://localhost:8080/api/users`
* > DELETE`http://localhost:8080/api/users/id`

## Body for POST and PUT endpoints:
```
{
    "id": "f05667e0-8f32-4566-a318-5fbbc8baaaf3",
    "name": "Adam",
    "address": "test 123"
}
```
For post requests you can omit the id.
