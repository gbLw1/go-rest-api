# Go REST API

This project is a secure REST API with JWT authentication, storing access tokens in cookies, and protecting sensitive routes through middleware.

---

## Dependencies

- [Air](https://github.com/cosmtrek/air) - Hot reload
- [GoDotEnv](https://github.com/joho/godotenv) - Environment variables
- [Gin Web Framework](https://gin-gonic.com/) - Web framework
- [Gorm](https://gorm.io/) - Database ORM
- [Bcrypt](https://github.com/golang/crypto) - Password hashing
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token

---

## Database

Tried with SQL Server but can't connect to port 1433 using TelNet so I use [PostgreSQL](https://www.elephantsql.com/)

issue: <https://github.com/hashicorp/vault/issues/4167>

---

## Setting up the project

Follow the steps below to set up the project.

### Environment Variables

create a `.env` file in the root directory and add the following:

```sh
PORT=8080
DB_URL="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
SECRET="Your_Super_Secret_Key"
```

`PORT` is the port where the API will run (default: 8080). You can change it if you want.

`DB_URL` is the connection string for the database. Change it according to your database credentials.

`SECRET` is the secret key for JWT authentication.

---

## Run

Run the following command to start the server:

```sh
go run ./cmd/main.go
```

### Running with Air (optional)

Air is a tool for running Go applications in the background, refreshing whenever it is modified.

If you got `air` installed, run the following command

```sh
air
```

---

## Testing

You can use [Postman](https://www.postman.com/) to test the API.

Base address: `http://localhost:{port}/`

ps: Check the port in `.env` file and feel free to change it.

---

## Endpoints

- [x] [POST] `~/signup`
- [x] [POST] `~/login`
- [x] [GET] `~/posts`
- [x] [GET] `~/posts/{id}`
- [x] [POST] `~/posts` (Requires authentication)
- [x] [PUT] `~/posts/{id}` (Requires authentication)
- [x] [DELETE] `~/posts/{id}` (Requires authentication)
