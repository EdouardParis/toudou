# Toudou: a mini tasks manager api

## Set up your api

1. Create PostgreSQL Database:
  _ install Postgre
  _ create database and user
  _ import database set up with 
    ```toudou=# \i database.sql; ```

2. Get dependencies and compile

```
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm

```

## Test you api


```
curl -i -X POST -H "Content-Type:application/json" -d "{ \"name\": \"third task\", \"description\":\"find some coffee\"}" http://localhost:8080/tasks

```

```
curl -i -X PATCH -H "Content-Type:application/json" -d "{ \"name\": \"task 1 updated\", \"description\":\"pants are optionnal\"}" http://localhost:8080/tasks/1

```
