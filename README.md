# Go REST API

A simple REST API written in Go.

---

## Dependencies

- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
- [GoDotEnv](https://github.com/joho/godotenv)
- [Gin Web Framework](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)

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

### Database Migration

Run the following command in the root directory:

```sh
go run .\migrate\migrate.go
```

---

## Running the API with CompileDaemon

CompileDaemon is a tool for running Go applications in the background, refreshing whenever it is modified.

Run the `CompileDaemon` command pointing to the module name in the [`go.mod`](./go.mod) file

```sh
CompileDaemon -command="./go-rest-api"
```

You can see the output in the terminal and you can kill the process with `Ctrl + C` at any time.

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
