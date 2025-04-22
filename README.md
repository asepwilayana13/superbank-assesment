# GOLANG FIBER

This is a RESTful API service built with Go-FIBER.

#### Folder
 | superbank-assesment
    | document [user guide aplikasi](document/Readme.md)
    | frontend
    | backend

### BACKEND

## Prerequisites

- [Go]  (https://go.dev/dl/)
- [FIBER] (https://gofiber.io/) - GO WEB FRAMEWORK
- [GORM]  (https://gorm.io/) - ORM (Object-Relational Mapping)
- [POSTGRES]  (https://www.postgresql.org/download/) - DATABASE
- [AIR] (go install github.com/cosmtrek/air@latest) - Live Reloading untuk Development

### Docker Setup (Recommended)

1. Step
   ```shell
    docker-compose up --build -d
    ```
2. URL BACKEND

    ```shell
    http://localhost:8000
    ```
3. URL FRONTEND

    ```shell
    http://localhost:3000/login
    ```
4. Halaman login admin url (http://localhost:3000/login)
   - masukan username dan password
    ```json
    {
      username: admin@example.com,
      password: today1234
    }
    ```
## Manual Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/asepwilayana13/superbank-assesment.git
    ```
## BACKEND
1. change directory
    ```shell
   cd backend
    ```

2. Install the dependencies:

    ```shell
    go mod tidy
    ```

3. Set up the database:

    - Create a postgres database named `example`.
    - Rename `.env.example` to `.env`
    - update the database connection details in the `.env` file.

4. Build and run the application:

    ```shell
    go run main.go
    ```
    atau

    ```shell
    air
    ```
### FRONTEND
## Installation
1. change directory
    ```shell
   cd frontend
    ```
2. Install the dependencies:

    ```shell
    npm install
    ```
3. Build and run the application:

    ```shell
    npm run dev
    ```

## UNIT TEST
## FOR LINUX
  Generate Unit Test:
  - change folder backend
    ```shell
      cd backend
    ```
  - running unit test
    ```shell
      ./unitTest.sh
    ```

## FOR WINDOWS
Generate Unit Test:
  - change folder backend
    ```shell
      cd backend
    ```
  - generate coverage
    ```shell
     make test
    ```
  - duplicate
    ```shell
     make duplicate
    ```
  - filter
    ```shell
     make filter
    ```
  - coverage per fungsi
    ```shell
    make cover
    ```
  - Generate HTML coverage
    ```shell
    make cover-html
     ```


## API Endpoints
- `/v1/auth/login` - POST request to authenticate a user.
- `/v1/auth/register` - POST request to register a new user.
- comming soon

## SWAGGER

- [Url] http://localhost:3000/swagger/index.html
- ['Generate'] swag init -g main.go --output docs

## DOCS API

- [Url Document Api](document/myapp-go.postman_collection.json)



