# GOLANG FIBER

This is a RESTful API service built with Go-FIBER.

#### Folder
 | superbank-assesment
    | document
    | froentend
    | backend

### BACKEND

## Prerequisites

- [Go]  (https://go.dev/dl/)
- [FIBER] (https://gofiber.io/) - GO WEB FRAMEWORK
- [GORM]  (https://gorm.io/) - ORM (Object-Relational Mapping)
- [POSTGRES]  (https://www.postgresql.org/download/) - DATABASE
- [AIR] (go install github.com/cosmtrek/air@latest) - Live Reloading untuk Development

## Installation

1. Clone the repository:

    ```shell
    git clone https://gitlab.smooets.com/aswil/superbank-assesment.git
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
5. Generate Unit Test:
    - generate coverage
        ```shell
         make test
        ```
    - coverage per fungsi
         ```shell
         make test
        ```
    - Buka tampilan HTML coverage
         ```shell
         make test
        ```


## API Endpoints
- `/v1/auth/login` - POST request to authenticate a user.
- `/v1/auth/register` - POST request to register a new user.
- comming soon

## SWAGGER

- [Url] http://localhost:3000/swagger/index.html
- ['Generate'] swag init -g main.go --output docs



### FRONTEND
## Installation

1. Clone the repository:
   ```shell
    git clone https://gitlab.smooets.com/aswil/superbank-assesment.git
    ```
2. Install the dependencies:

    ```shell
    npm insatall
    ```
3. Build and run the application:

    ```shell
    npm run dev
    ```

### Docker Setup (Recommended)

1. Step
   ```shell
    docker-compose up --build -d
    ```
2. URL BACKEND

    ```shell
    http://localhost:8000
    ```
3. URL ENV

    ```shell
    http://localhost:3000
    ```

