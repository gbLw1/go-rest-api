# Go REST API

A simple REST API written in Go.

---

## Dependencies

- [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Hot reload
- [GoDotEnv](https://github.com/joho/godotenv) - Environment variables
- [Gin Web Framework](https://gin-gonic.com/) - Web framework
- [Gorm](https://gorm.io/) - Database ORM

---

## Database

Tried with SQL Server but can't connect to 1433 using TelNet so I use [PostgreSQL](https://www.elephantsql.com/).

issue: <https://github.com/hashicorp/vault/issues/4167>

---

## Setting up the project

Follow the steps below to set up the project.

### Environment Variables

create a `.env` file in the root directory and add the following:

```sh
PORT=8080
DB_URL="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
```

`PORT` is the port where the API will run (default: 8080). You can change it if you want.

`DB_URL` is the connection string for the database. Change it according to your database credentials.

---

## Run

Run the following command to start the server:

```sh
go run main.go
```

---

## Testing

You can use [Postman](https://www.postman.com/) to test the API.

Base address: `http://localhost:{port}/`

ps: Check the port in `.env` file and feel free to change it.

---

## Endpoints

- [x] [GET] `~/posts`
- [x] [GET] `~/posts/{id}`
- [x] [POST] `~/posts`
- [x] [PUT] `~/posts/{id}`
- [x] [DELETE] `~/posts/{id}`
