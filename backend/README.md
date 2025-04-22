# GOLANG FIBER

This is a RESTful API service built with Go-FIBER.

## Prerequisites

- [Go]  (https://go.dev/dl/)
- [FIBER] (https://gofiber.io/) - GO WEB FRAMEWORK
- [GORM]  (https://gorm.io/) - ORM (Object-Relational Mapping)
- [POSTGRES]  (https://www.postgresql.org/download/) - DATABASE
- [AIR] (go install github.com/cosmtrek/air@latest) - Live Reloading untuk Development

## Installation

1. Clone the repository:

    ```shell
    git clone https://gitlab.smooets.com/aswil/boilerplate-go-gorm-fiber.git
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

For detailed information about each endpoint and the expected request/response formats, please refer to the [API documentation](./docs/api.md).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](./LICENSE).
