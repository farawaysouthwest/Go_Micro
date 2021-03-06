# Go_Micro

A minimalist Go microservice architecture example. Use as a template for your next MicroService!

## Getting Started

Uses basic Go commands for local development.

Install packages

```sh
go install
```

Run Locally

```sh
go run .
```

### Prerequisites

- Go sdk
- Mockery testing mock generator
- Docker for deployment

## Project Structure

### Controller

### Service

### Model

## Running the tests

Testing mocks can be generated with the Mockery tool, this will generate mocks for any public interface.

```sh
mockery --all
```

Then just run all tests:

```sh
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
