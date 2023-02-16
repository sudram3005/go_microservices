# Simple Go Micro Service Project

This Project mainly contains four micro services written in golang communincating using rest API 

## Tech used
1. Golang for writing the microservices
2. [Gin](https://pkg.go.dev/github.com/gin-gonic/gin) for routing purpose
3. [Sqlx](https://github.com/jmoiron/sqlx) for db related operations
4. Postgres for data persistence

## Micro-services
1. **Authentication service:** This service authenticates user for access. Takes users email and password and authenticate them
2. **Employee-service:** This service do the CRUD operations on postgres db.
3. **Broker-service:** This can communicates with all other services.
4. **Front-end service:** This is the frontend of the microservice

## Setup local development



### Setup infrastructure

In [project](./project/) folder there is a Makefile where you can build and start individual docker-services. For starting and running all microservices as docker images use
``` bash
    make up
``` 
