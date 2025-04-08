###### Install the package to be able to run migrations
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

###### Run migrations
    migrate -database "postgres://postgres:postgres@localhost:5432/gateway?sslmode=disable" -path migrations up

###### Run app
    go run cmd/app/main.go

### 🛠️ Stack

![REST](https://img.shields.io/badge/REST-API-000000?style=for-the-badge&logo=rest&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff&style=for-the-badge)